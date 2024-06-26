package json

import (
	"io"
	"net/http"
	"net/url"
	"sql_filler/subjects/assignment"
	"sql_filler/webAPI/user"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func (bec *backEnd) serveCompleated(w http.ResponseWriter, r *http.Request) {
	userID, err := user.CheckCookieAndGetUserId(bec.db, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Errorf("lesson serveCompleated body read error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	q, err := url.ParseQuery(string(bytes))
	if err != nil {
		log.Errorf("lesson serveCompleated Pasrse error: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for k, v := range q {
		if strings.Compare("keys[]", k) == 0 {
			for _, idString := range v {
				id, err := strconv.Atoi(idString)
				if err != nil {
					log.Errorf("lesson serveCompleated atoi error: %s", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				err = assignment.StartAssignment(bec.db, userID, id)
				if err != nil {
					log.Errorf("lesson serveCompleated StartAssignment error: %s", err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
		}
	}
	w.WriteHeader(http.StatusOK)
}
