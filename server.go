package main

import (
	_ "fmt"
	"html/template"
	"net/http"
)

type Person struct {
	Name string
}

var t = template.Must(template.ParseFiles("index.html"))

func IndexHandler(w http.ResponseWriter,
	r *http.Request) {

	person := Person{
		Name: r.URL.Query().Get("name"),
	}

	t.Execute(w, person)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":3000", nil)
}
