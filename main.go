package main

import (
	"fmt"
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

func setupRoutes(db *godb.DB) *mux.Router {
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

	router.PathPrefix("/user/{function}").Handler(http.StripPrefix("/user/", my_user.NewBackEnd(db)))
	router.PathPrefix("/review/").Handler(http.StripPrefix("/review/", review.NewBackEnd(db)))
	router.PathPrefix("/lesson/").Handler(http.StripPrefix("/lesson/", lesson.NewBackEnd(db)))
	router.PathPrefix("/json/{function}").Handler(http.StripPrefix("/json/", json.NewBackEnd(db)))
	router.PathPrefix("/learning/{function}").Handler(http.StripPrefix("/learning/", learining.NewBackEnd(db)))

	return router
}

func cacheControlWrapper(h http.Handler, hours int) http.Handler {
	seconds := hours * 3600
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", seconds)) // 1 day
		h.ServeHTTP(w, r)
	})
}
func serveStaticFiles(router *mux.Router) {
	frontFiles := http.FileServer(http.Dir("./front/"))
	router.PathPrefix("/front/").Handler(
		http.StripPrefix("/front/", cacheControlWrapper(frontFiles, 12)),
	)

	staticDir := "./" + subject.Static_file_dir_name + "/"
	staticFiles := http.FileServer(http.Dir(staticDir))
	router.PathPrefix("/" + subject.Static_file_dir_name + "/").Handler(
		http.StripPrefix("/"+subject.Static_file_dir_name+"/", cacheControlWrapper(staticFiles, 7*24)),
	)
}

func main() {
	logger.SetUp(log.InfoLevel)

	db, err := godb.Open(sqlite.Adapter, "./subjects.db")
	if err != nil {
		log.Fatal("Error Opening database :", err)
	}

	router := setupRoutes(db)
	serveStaticFiles(router)

	log.Warn("Starting new server")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		db.Close()
		log.Fatal("Error Starting the HTTP Server :", err)
	}

	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
