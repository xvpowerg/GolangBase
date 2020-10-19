package main

import (
	"html/template"
	"log"
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
	user1 := User{
		ID:   10,
		Name: "Ken",
	}
	user2 := User{
		ID:   20,
		Name: "Vivin",
	}
	user3 := User{
		ID:   30,
		Name: "Lindy",
	}

	userList := []User{user1, user2, user3}
	err := temp.Execute(res, userList)
	if err != nil {
		log.Println("Error!", err)
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(home))
	http.ListenAndServe(":8080", nil)
}
