package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter,
	r *http.Request) {

	name := r.URL.Query().Get("name")

	fmt.Fprint(w, "<h1>hello "+name+" !</h1>")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":3000", nil)
}
