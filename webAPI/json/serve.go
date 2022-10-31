package json

import (
	"encoding/json"
	"net/http"
	"path"
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
	body := json.NewDecoder(r.Body)
	var progress map[string][]interface{}
	err := body.Decode(&progress)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = bec.saveProgress(101, progress)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (bec *backEnd) serveRadical(w http.ResponseWriter, r *http.Request) {
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

	serveBodyJson(w, radical)
}

func (bec *backEnd) serveKanji(w http.ResponseWriter, r *http.Request) {
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

	serveBodyJson(w, kanji)
}

func (bec *backEnd) serveVocabulary(w http.ResponseWriter, r *http.Request) {
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

	serveBodyJson(w, vocabulary)
}

func serveBodyJson(w http.ResponseWriter, jsonStr interface{}) {
	buf, err := json.Marshal(jsonStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}
