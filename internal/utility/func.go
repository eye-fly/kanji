package utility

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func ServeBodyJson(w http.ResponseWriter, jsonStr interface{}) {
	buf, err := json.Marshal(jsonStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Errorf("marshal error: %s", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}
