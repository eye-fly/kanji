package json

import (
	"encoding/json"
	"net/http"

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

	path := r.URL.Path
	if path == "progress" {
		bec.serveProgress(w, r)

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

	err = bec.saveProgress(0, progress)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (bec *backEnd) serveKanji(w http.ResponseWriter, r *http.Request) {

}
