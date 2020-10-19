package main

import (
	"io"
	"net/http"
)

var html string = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>MyFistWeb</title>
</head>
<body>
    <p>Home!GGGXXXX</p>
</body>
</html>`

var lgoinHTML string = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>MyFistWeb</title>
</head>
<body>
    <p>Login!!</p>
</body>
</html>`

func home(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, html)
}
func login(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, lgoinHTML)
}
func main() {
	// http.Handle("/", http.HandlerFunc(home))
	// http.Handle("/login", http.HandlerFunc(login))
	http.ListenAndServe(":8080", nil)
}
