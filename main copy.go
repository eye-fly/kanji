package main

import (
	"net/http"
	"sql_filler/waniAPI/subject"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// fronReview "sql_filler/front/review"

func main() {
	// c := &http.Client{}
	// v, err := subject.RequestVocabulary(c, map[string]string{"ids": "9182"})

	// println(0)
	// if err != nil {
	// 	println("err", err.Error())
	// } else {
	// 	println(v[0].URL)
	// }

	start := 21
	end := 25
	lvls := strconv.Itoa(start)
	for i := start + 1; i <= end; i++ {
		lvls += "," + strconv.Itoa(i)
	}
	println(lvls)

	subject.GetAudioUrlAndDownloadIt("7,8,9,10")

	router := mux.NewRouter()
	// reviewBec := review.NewBackEnd(db)
	// lessonBec := lesson.NewBackEnd(db)
	// jsonBec := json.NewBackEnd(db)
	// router.PathPrefix("/user/{function}").Handler(http.StripPrefix("/user/", userBec))
	// router.PathPrefix("/review/{function}").Handler(http.StripPrefix("/review/", reviewBec))
	// router.PathPrefix("/lesson/{function}").Handler(http.StripPrefix("/lesson/", lessonBec))
	// router.PathPrefix("/json/{function}").Handler(http.StripPrefix("/json/", jsonBec))

	front := http.FileServer(http.Dir("./" + subject.Static_file_dir_name + "/"))
	router.PathPrefix("/" + subject.Static_file_dir_name + "/").Handler(http.StripPrefix("/"+subject.Static_file_dir_name+"/", front))

	log.Warn("Starting new server")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server :", err)
		return
	}

}
