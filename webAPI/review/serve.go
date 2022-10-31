package review

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sql_filler/subjects"
	"strconv"
	"strings"

	"github.com/samonzeweb/godb"
	log "github.com/sirupsen/logrus"
)

type backEnd struct {
	db *godb.DB
}

func NewBackEnd(db *godb.DB) *backEnd {
	return &backEnd{
		db: db,
	}
}

func (bec *backEnd) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	if path == "queue" {
		bec.serveQueue(w, r)
	} else if path == "items" {
		bec.serveItems(w, r)
	} else {
		w.WriteHeader(http.StatusBadGateway)
	}

}

func (bc *backEnd) serveQueue(w http.ResponseWriter, r *http.Request) {
	queue, err := GetQueue(bc.db, 101)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("GetQueue: %s", err)
		return
	}
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, "[")
	for i, v := range queue {
		fmt.Fprint(w, v)
		if i < len(queue)-1 {
			fmt.Fprint(w, ",")
		}
	}
	fmt.Fprint(w, "]")
}

func (bc *backEnd) serveItems(w http.ResponseWriter, r *http.Request) {
	s := strings.Split(r.URL.Query().Get("ids"), ",")
	ids := make([]int, len(s))
	var err error
	for i, v := range s {
		ids[i], err = strconv.Atoi(v)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Errorf("BadRequest ids: %s", s)
			return
		}
	}

	buf := []byte{}
	buf = append(buf, []byte("[")...)
	// fmt.Fprint(buf)
	for i, id := range ids {
		typ, err := subjects.GetType(bc.db, id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			log.Errorf("GetType: %s", err)
			return
		}

		var revieItem interface{}
		if typ == subjects.TypeRadical {
			revieItem, err = GetRadical(bc.db, id, 101)
		} else if typ == subjects.TypeKanji {
			revieItem, err = GetKanji(bc.db, id, 101)
		} else if typ == subjects.TypeVocabulary {
			revieItem, err = GetVocabulary(bc.db, id, 101)
		} else {
			err = fmt.Errorf("no maching type")
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			log.Errorf("GetSubject: %s", err)
			return
		}

		bytes, err := json.Marshal(revieItem)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			log.Errorf("Marshal: %s", err)
			return
		}
		buf = append(buf, bytes...)
		if i < len(ids)-1 {
			buf = append(buf, []byte(",")...)
		}
	}
	buf = append(buf, []byte("]")...)

	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}
