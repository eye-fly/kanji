package main

import (
	"net/http"

	// fronReview "sql_filler/front/review"
	"sql_filler/front"
	"sql_filler/internal/logger"
	"sql_filler/waniAPI/subject"
	"sql_filler/webAPI/json"
	"sql_filler/webAPI/learining"
	"sql_filler/webAPI/lesson"
	"sql_filler/webAPI/review"
	my_user "sql_filler/webAPI/user"

	"github.com/gorilla/mux"
	"github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/sqlite"
	log "github.com/sirupsen/logrus"
)

func main() {
	logger.SetUp(log.InfoLevel)

	db, err := godb.Open(sqlite.Adapter, "./subjects.db")
	panicIfErr(err, db)

	// b, err := users.CanBeUnlocked(db, 101, 500)
	// panicIfErr(err, db)
	// fmt.Println(b)

	// for i := 40; i <= 60; i += 5 {
	// err = subject.GetAndPutAllSubjects(db, fmt.Sprintf("%v,%v,%v,%v,%v", i, i+1, i+2, i+3, i+4))
	// 	panicIfErr(err, db)
	// 	fmt.Println(fmt.Sprint(i))
	// 	time.Sleep(time.Second * 20)
	// }

	// client := &http.Client{}
	// colection, err := subject.RequestAssigment(client, map[string]string{
	// 	"levels": "1,2,3,4,5",
	// })
	// panicIfErr(err, db)
	// for _, subjct := range colection {
	// 	subjct.UserId = 101
	// 	err = subjct.AddToDB(db, "y")
	// 	panicIfErr(err, db)
	// }

	fnt := front.NewBackEnd(db)

	router := mux.NewRouter()
	router.StrictSlash(true)
	router.HandleFunc("/", fnt.MainPageHandler)
	router.HandleFunc("/user/login", front.LoginHandler)
	router.HandleFunc("/user/register", fnt.RegisterHandler)
	router.HandleFunc("/review/session", fnt.ReviewHandler)
	router.HandleFunc("/review/learning/session", fnt.ReviewLearningHandler)
	router.HandleFunc("/lesson/session", fnt.SesionHandler)
	router.HandleFunc("/lesson/learning/session", fnt.LessonLerningHandler)
	router.HandleFunc("/learning/find", fnt.LearningFindAdderHandler)

	userBec := my_user.NewBackEnd(db)
	reviewBec := review.NewBackEnd(db)
	lessonBec := lesson.NewBackEnd(db)
	jsonBec := json.NewBackEnd(db)
	learningBec := learining.NewBackEnd(db)
	router.PathPrefix("/user/{function}").Handler(http.StripPrefix("/user/", userBec))
	router.PathPrefix("/review/").Handler(http.StripPrefix("/review/", reviewBec))
	router.PathPrefix("/lesson/").Handler(http.StripPrefix("/lesson/", lessonBec))
	router.PathPrefix("/json/{function}").Handler(http.StripPrefix("/json/", jsonBec))
	router.PathPrefix("/learning/{function}").Handler(http.StripPrefix("/learning/", learningBec))

	front := http.FileServer(http.Dir("./front/"))
	router.PathPrefix("/front/").Handler(http.StripPrefix("/front/", front))
	files := http.FileServer(http.Dir("./" + subject.Static_file_dir_name + "/"))
	router.PathPrefix("/" + subject.Static_file_dir_name + "/").Handler(http.StripPrefix("/"+subject.Static_file_dir_name+"/", files))

	log.Warn("Starting new server")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server :", err)
		return
	}

	err = db.Close()
	panicIfErr(err, nil)
}

func panicIfErr(err error, db *godb.DB) {
	if err != nil {
		db.Close()
		log.Fatal(err)
	}
}
