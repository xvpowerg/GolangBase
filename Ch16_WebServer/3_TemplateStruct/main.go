package main

import (
	"html/template"
	"net/http"
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
	user := User{
		ID:   10,
		Name: "Ken",
	}
	temp.Execute(res, user)
}

func main() {
	http.Handle("/", http.HandlerFunc(home))
	http.ListenAndServe(":8080", nil)
}
