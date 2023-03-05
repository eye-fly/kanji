package front

import (
	"html/template"
	"net/http"
)

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
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/register/"):]
	p := &Page{Title: title}
	t, _ := template.ParseFiles("front/user/register/register.html")
	t.Execute(w, p)
}


func ReviewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/session/"):]
	p := &Page{Title: title}
	t, _ := template.ParseFiles("front/review/session.html")
	t.Execute(w, p)
}