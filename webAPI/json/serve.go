package json

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"

	"github.com/samonzeweb/godb"
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
	b, err := path.Match("progress", p)
	if err != nil {
		fmt.Printf("json serveHTTP error: %s", err)
	}
	if b {
		bec.serveProgress(w, r)
	}
	b, err = path.Match("lesson/completed", p)
	if err != nil {
		fmt.Printf("json serveHTTP error: %s", err)
	}
	if b {
		bec.serveCompleated(w, r)
	}

	b, err = path.Match("radical/*", p)
	if err != nil {
		fmt.Printf("json serveHTTP error: %s", err)
	}
	if b {
		bec.serveRadical(w, r)
	}
	b, err = path.Match("kanji/*", p)
	if err != nil {
		fmt.Printf("json serveHTTP error: %s", err)
	}
	if b {
		bec.serveKanji(w, r)
	}
	b, err = path.Match("vocabulary/*", p)
	if err != nil {
		fmt.Printf("json serveHTTP error: %s", err)
	}
	if b {
		bec.serveVocabulary(w, r)
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
