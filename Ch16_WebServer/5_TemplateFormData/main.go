package main

import (
	"html/template"
	"net/http"
	"strconv"
)

var temp *template.Template

type User struct {
	ID   int
	Name string
}

func init() {

	temp = template.Must(template.ParseFiles("index.html"))
}

func home(res http.ResponseWriter, req *http.Request) {
	idstr := req.FormValue("myId")
	name := req.FormValue("myName")
	id, ok := strconv.ParseInt(idstr, 10, 32)
	if ok != nil {
		id = 0
	}
	user := User{
		ID:   int(id),
		Name: name,
	}
	temp.Execute(res, user)
}

func main() {
	http.Handle("/", http.HandlerFunc(home))
	http.ListenAndServe(":8080", nil)
}
