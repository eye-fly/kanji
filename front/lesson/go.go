package front

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func SesionHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/session/"):]
	p := &Page{Title: title}
	t, _ := template.ParseFiles("front/lesson/session.html")
	t.Execute(w, p)
}
