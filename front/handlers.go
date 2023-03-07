package front

import (
	"html/template"
	"net/http"
	"sql_filler/webAPI/user"

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

type Page struct {
	Title string
	Body  []byte
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/login/"):]
	p := &Page{Title: title}
	t, _ := template.ParseFiles("front/user/login/login.html")
	t.Execute(w, p)
}

func (bec *backEnd) MainPageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := user.CheckCookieAndGetUserId(bec.db, r)
	if err != nil {
		http.Redirect(w, r, "/user/login?from=/", http.StatusFound)
	}
	serveHTML(w, "mainpage", "front/mainpage/main.html")
}

func (bec *backEnd) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	serveHTML(w, "register", "front/mainpage/main.html")
	// title := r.URL.Path[len("/register/"):]
	// p := &Page{Title: title}
	// t, _ := template.ParseFiles("front/user/register/register.html")
	// t.Execute(w, p)
}

func (bec *backEnd) SesionHandler(w http.ResponseWriter, r *http.Request) {
	serveHTML(w, "lesson", "front/lesson/session.html")
	// title := r.URL.Path[len("/session/"):]
	// p := &Page{Title: title}
	// t, _ := template.ParseFiles("front/lesson/session.html")
	// t.Execute(w, p)
}

func (bec *backEnd) ReviewHandler(w http.ResponseWriter, r *http.Request) {
	serveHTML(w, "review", "front/review/session.html")
	// title := r.URL.Path[len("/session/"):]
	// p := &Page{Title: title}
	// t, _ := template.ParseFiles("front/review/session.html")
	// t.Execute(w, p)
}

func serveHTML(w http.ResponseWriter, title, htmlPath string) {
	p := &Page{Title: title}
	t, e := template.ParseFiles(htmlPath)
	if e != nil {
		w.Write([]byte(e.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	t.Execute(w, p)
}
