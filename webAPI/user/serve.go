package lesson

import (
	"encoding/json"
	"net/http"
	"path"
	"sql_filler/subjects/users"

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

	p := r.URL.Path

	b, err := path.Match("queue", p)
	if err != nil {
		log.Errorf("json serveHTTP error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if b {
		bec.serveQueue(w, r)
	}

}

func (bec *backEnd) serveQueue(w http.ResponseWriter, r *http.Request) {
	userID := 101
	err := users.CanLevelup(bec.db, userID)
	if err != nil {
		log.Errorf("lesson serveQueue CanLevelup error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = users.UnlockLockedSubject(bec.db, 101)
	if err != nil {
		log.Errorf("lesson serveQueue UnlockLockedSubject error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	q, err := getLessonQueue(bec.db, userID)
	if err != nil {
		log.Errorf("lesson serveQueue getLessonQueue error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	serveBodyJson(w, q)
}

func serveBodyJson(w http.ResponseWriter, jsonStr interface{}) {
	buf, err := json.Marshal(jsonStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Errorf("lesson serveQueue marshal error: %s", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}