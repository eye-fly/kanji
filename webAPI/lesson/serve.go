package lesson

import (
	"net/http"
	"path"
	"sql_filler/internal/utility"
	"sql_filler/subjects/users"
	"sql_filler/webAPI/user"

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
		{"queue", bec.serveQueue},
		{"queue_count", bec.serveQueueCount},
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

func (bec *backEnd) serveQueue(w http.ResponseWriter, r *http.Request) {
	userID, err := user.CheckCookieAndGetUserId(bec.db, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = users.CanLevelup(bec.db, userID)
	if err != nil {
		log.Errorf("lesson serveQueue CanLevelup error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = users.UnlockLockedSubject(bec.db, userID)
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
	utility.ServeBodyJson(w, q)
}

func (bec *backEnd) serveQueueCount(w http.ResponseWriter, r *http.Request) {
	userID, err := user.CheckCookieAndGetUserId(bec.db, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = users.CanLevelup(bec.db, userID)
	if err != nil {
		log.Errorf("lesson serveQueue CanLevelup error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = users.UnlockLockedSubject(bec.db, userID)
	if err != nil {
		log.Errorf("lesson serveQueue UnlockLockedSubject error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ids, err := getLessonQueueId(bec.db, userID)
	if err != nil {
		log.Errorf("lesson serveQueue getLessonQueue error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var qc queueCountJson
	qc.Count = len(ids)

	utility.ServeBodyJson(w, &qc)
}
