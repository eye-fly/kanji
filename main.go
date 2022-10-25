package main

import (
	"net/http"
	fronReview "sql_filler/front/review"
	"sql_filler/webAPI/json"
	"sql_filler/webAPI/review"

	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/sqlite"
)

func main() {

	db, err := godb.Open(sqlite.Adapter, "./subjects.db")
	panicIfErr(err, db)
	// OPTIONAL: Set logger to show SQL execution logs
	db.SetLogger(log.New(os.Stderr, "", 0))

	// err = subject.GetAndPutAllSubjects(db, "2,3")
	// if err != nil {
	// 	panic(err)
	// }

	// client := &http.Client{}
	// colection, err := subject.RequestAssigment(client, map[string]string{
	// 	"levels": "1,2,3",
	// })
	// panicIfErr(err, db)
	// for _, subjct := range colection {
	// 	subjct.UserId = 101
	// 	err = subjct.AddToDB(db, "y")
	// 	panicIfErr(err, db)
	// }

	router := mux.NewRouter()
	router.HandleFunc("/review/session", fronReview.SesionHandler)
	reviewBec := review.NewBackEnd(db)
	router.PathPrefix("/review/{function}").Handler(http.StripPrefix("/review/", reviewBec))
	jsonBec := json.NewBackEnd(db)
	router.PathPrefix("/json/{function}").Handler(http.StripPrefix("/json/", jsonBec))
	// router.HandleFunc("/json/progress", review.UpdateItem).Methods(http.MethodPut)
	// styles := http.FileServer(http.Dir("./review/"))
	// router.PathPrefix("/review/").Handler(http.StripPrefix("/review/", styles))

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
