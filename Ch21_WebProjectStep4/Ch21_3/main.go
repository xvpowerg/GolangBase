package main

import (
	"html/template"
	"log"
	"net/http"

	"tw.com.maskweb/obj"
	"tw.com.maskweb/tools"
)

var tpl *template.Template
var positionList []*obj.Position

func init() {
	tpl = template.Must(template.ParseFiles("../template/main.html"))
	positionList = tools.GetPositionList()
}

func home(res http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(res, nil)
	if err != nil {
		log.Println("err:", err)
	}
}

func main() {
	//註冊 根目錄網址
	http.Handle("/", http.HandlerFunc(home))
	http.ListenAndServe(":8080", nil)
}
