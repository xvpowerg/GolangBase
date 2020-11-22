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

func addr(res http.ResponseWriter, req *http.Request) {
	addr := req.FormValue("addr")
	//產生maskCount
	maskList := tools.LoadingMaskList(positionList, nil)
	//依地址過濾
	result := tools.SearchByAddress(addr, maskList)

	err := tpl.Execute(res, result)
	if err != nil {
		log.Println("err:", err)
	}
}

func main() {
	//註冊網址為/addr
	http.Handle("/addr", http.HandlerFunc(addr))
	//註冊 根目錄網址
	http.Handle("/", http.HandlerFunc(home))

	http.ListenAndServe(":8080", nil)
}
