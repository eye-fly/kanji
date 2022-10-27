package lesson

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"

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
	// b, err := path.Match("progress", p)
	// if err != nil {
	// 	fmt.Printf("json serveHTTP error: %s", err)
	// }
	// if b {
	// 	bec.serveProgress(w, r)
	// }

	b, err := path.Match("queue", p)
	if err != nil {
		fmt.Printf("json serveHTTP error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if b {
		bec.serveQueue(w, r)
	}

}

func (bec *backEnd) serveQueue(w http.ResponseWriter, r *http.Request) {
	q, err := getLessonQueue(bec.db, 101)
	if err != nil {
		fmt.Printf("lesson serveQueue error: %s", err)
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
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}
