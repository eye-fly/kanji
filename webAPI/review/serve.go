package review

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"sql_filler/internal/utility"
	"sql_filler/subjects"
	"sql_filler/subjects/users"
	"sql_filler/webAPI/user"
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

type serves struct {
	path       string
	handleFunc func(w http.ResponseWriter, r *http.Request)
}

func (bec *backEnd) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	serves := []serves{
		{"learning/queue", bec.serveQueueLearningOnly},
		{"queue", bec.serveQueue},
		{"queue_count", bec.serveQueueCount},
		{"learning/queue_count", bec.serveQueueLearningCount},
		{"items", bec.serveItems},
		{"", redirect},
	}

	p := r.URL.Path
	for _, srv := range serves {
		b, err := path.Match(srv.path, p)
		if err != nil {
			log.Errorf("json serveHTTP error: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if b {
			srv.handleFunc(w, r)
			return
		}
	}
	w.WriteHeader(404)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
func (bc *backEnd) serveQueueOpt(w http.ResponseWriter, r *http.Request, onlyLearning bool) {
	userID, err := user.CheckCookieAndGetUserId(bc.db, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	queue, err := GetQueueOpt(bc.db, userID, onlyLearning)
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
func (bc *backEnd) serveQueue(w http.ResponseWriter, r *http.Request) {
	bc.serveQueueOpt(w, r, false)
}
func (bc *backEnd) serveQueueLearningOnly(w http.ResponseWriter, r *http.Request) {
	bc.serveQueueOpt(w, r, true)
}

type queueCountJson struct {
	Count int `json:"count"`
}

func (bec *backEnd) serveQueueCount(w http.ResponseWriter, r *http.Request) {
	bec.serveQueueCountOpt(w, r, false)
}
func (bec *backEnd) serveQueueLearningCount(w http.ResponseWriter, r *http.Request) {
	bec.serveQueueCountOpt(w, r, true)
}
func (bec *backEnd) serveQueueCountOpt(w http.ResponseWriter, r *http.Request, LeariningOnly bool) {
	userID, err := user.CheckCookieAndGetUserId(bec.db, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	queue, err := GetQueueOpt(bec.db, userID, LeariningOnly)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("GetQueue: %s", err)
		return
	}

	var qc queueCountJson
	qc.Count = len(queue)

	utility.ServeBodyJson(w, &qc)
}

func (bc *backEnd) serveItems(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(users.CookieSesionIdName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sesion, err := users.IsSessionIdOk(bc.db, cookie.Value)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("GetCheckCookie: %s", err)
		return
	}

	s := strings.Split(r.URL.Query().Get("ids"), ",")
	ids := make([]int, len(s))
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
			revieItem, err = GetRadical(bc.db, id, sesion.UserId)
		} else if typ == subjects.TypeKanji {
			revieItem, err = GetKanji(bc.db, id, sesion.UserId)
		} else if typ == subjects.TypeVocabulary {
			revieItem, err = GetVocabulary(bc.db, id, sesion.UserId)
		} else {
			err = fmt.Errorf("no maching type")
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			log.Errorf("GetSubject '%s': %s", typ, err)
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
