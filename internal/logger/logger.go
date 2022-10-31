package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func SetUp(lvl log.Level) {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.SetLevel(lvl)
	log.SetReportCaller(true)
}
