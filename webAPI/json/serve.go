package json

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path"
	"sql_filler/subjects/users"
	"strconv"

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
		{"progress", bec.serveProgress},
		{"lesson/completed", bec.serveCompleated},
		{"radical/*", bec.serveRadical},
		{"kanji/*", bec.serveKanji},
		{"vocabulary/*", bec.serveVocabulary},
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
		}
	}
}
func (bec *backEnd) serveProgress(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(users.CookieSesionIdName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	session, err := users.IsSessionIdOk(bec.db, cookie.Value)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("GetCheckCookie: %s", err)
		return
	}

	body := json.NewDecoder(r.Body)
	var progress map[string][]interface{}
	err = body.Decode(&progress)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = bec.saveProgress(session.UserId, progress)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (bec *backEnd) serveRadical(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(users.CookieSesionIdName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = users.IsSessionIdOk(bec.db, cookie.Value)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("GetCheckCookie: %s", err)
		return
	}

	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	radical, err := getRadical(bec.db, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ServeBodyJson(w, radical)
}

func (bec *backEnd) serveKanji(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(users.CookieSesionIdName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = users.IsSessionIdOk(bec.db, cookie.Value)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("GetCheckCookie: %s", err)
		return
	}

	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	kanji, err := getKanji(bec.db, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ServeBodyJson(w, kanji)
}

func (bec *backEnd) serveVocabulary(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(users.CookieSesionIdName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = users.IsSessionIdOk(bec.db, cookie.Value)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("GetCheckCookie: %s", err)
		return
	}

	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	vocabulary, err := getVocabulary(bec.db, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ServeBodyJson(w, vocabulary)
}

func ServeBodyJson(w http.ResponseWriter, jsonStr interface{}) {
	buf, err := json.Marshal(jsonStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}
