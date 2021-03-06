package main

import (
	"html/template"
	"net/http"
)

var temp *template.Template

func init() {
	//temp, _ = template.ParseFiles("index.html")
	//用於簡化ParseFiles 的err處理
	temp = template.Must(template.ParseFiles("index.html"))
}

func home(res http.ResponseWriter, req *http.Request) {

	temp.Execute(res, nil)
}

func main() {
	http.Handle("/", http.HandlerFunc(home))

	http.ListenAndServe(":8080", nil)
}
