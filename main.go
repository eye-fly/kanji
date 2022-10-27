package main

import (
	"net/http"
	frontLesson "sql_filler/front/lesson"
	fronReview "sql_filler/front/review"
	"sql_filler/waniAPI/subject"
	"sql_filler/webAPI/json"
	"sql_filler/webAPI/review"

	"log"

	"github.com/gorilla/mux"
	"github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/sqlite"
)

func main() {

	db, err := godb.Open(sqlite.Adapter, "./subjects.db")
	panicIfErr(err, db)
	// db.SetLogger(log.New(os.Stderr, "", 0))

	// for i := 40; i <= 60; i += 5 {
	// 	err = subject.GetAndPutAllSubjects(db, fmt.Sprintf("%v,%v,%v,%v,%v", i, i+1, i+2, i+3, i+4))
	// 	panicIfErr(err, db)
	// 	fmt.Println(fmt.Sprint(i))
	// 	time.Sleep(time.Second * 20)
	// }

	client := &http.Client{}
	colection, err := subject.RequestAssigment(client, map[string]string{
		"levels": "1,2,3,4,5",
	})
	panicIfErr(err, db)
	for _, subjct := range colection {
		subjct.UserId = 101
		err = subjct.AddToDB(db, "y")
		panicIfErr(err, db)
	}

	router := mux.NewRouter()
	router.HandleFunc("/review/session", fronReview.SesionHandler)
	router.HandleFunc("/lesson/session", frontLesson.SesionHandler)
	reviewBec := review.NewBackEnd(db)
	router.PathPrefix("/review/{function}").Handler(http.StripPrefix("/review/", reviewBec))
	jsonBec := json.NewBackEnd(db)
	router.PathPrefix("/json/{function}").Handler(http.StripPrefix("/json/", jsonBec))
	front := http.FileServer(http.Dir("./front/"))
	router.PathPrefix("/front/").Handler(http.StripPrefix("/front/", front))

	// //tmp
	// lessQ := http.FileServer(http.Dir("./front/lesson/"))
	// router.PathPrefix("/lesson/").Handler(http.StripPrefix("/lesson/", lessQ))

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server :", err)
		return
	}

	// tab, err := review.GetQueue(db, 101)
	// panicIfErr(err, db)

	// for _, v := range tab {
	// 	println(v)
	// }

	err = db.Close()
	panicIfErr(err, nil)
}

// It's just an example, what did you expect ? (never do that in real code)
func panicIfErr(err error, db *godb.DB) {
	if err != nil {
		db.Close()
		panic(err)
	}
}
