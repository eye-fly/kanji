package learining

import (
	"encoding/json"
	"net/http"
	"path"
	"sql_filler/subjects/assignment"
	"sql_filler/webAPI/user"
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
		{"find/submit", bec.serveFindSubmit},
		{"add", bec.serveAddToLearning},
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

func (bec *backEnd) serveAddToLearning(w http.ResponseWriter, r *http.Request) {
	userID, err := user.CheckCookieAndGetUserId(bec.db, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	type js struct {
		ID int `json:"id"`
	}
	body := json.NewDecoder(r.Body)
	var data js
	err = body.Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = assignment.AddAssigmentToLearing(bec.db, userID, data.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (bec *backEnd) serveFindSubmit(w http.ResponseWriter, r *http.Request) {
	userID, err := user.CheckCookieAndGetUserId(bec.db, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body := json.NewDecoder(r.Body)
	var data findSubmitData
	err = body.Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	data.SearchText = strings.TrimSpace(data.SearchText)

	if data.Type == findSubmitTypeVocabulary {
		bec.findVoc(userID, data, w)
		return
	} else if data.Type == findSubmitTypeKanji {
		bec.findKanji(userID, data, w)
		return
	}
}
