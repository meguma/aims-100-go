package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Id string
}

type Aim struct {
	Seq     int
	Content string
}

type Aims struct {
	Aims []Aim
}

func ApiHandler(w http.ResponseWriter,
	r *http.Request) {

	user := User{
		Id: r.URL.Query().Get("id"),
	}

	if user.Id == "1" {
		aim := Aim{
			Seq:     1,
			Content: "test",
		}

		aims := Aims{}

		aims.Aims = append(aims.Aims, aim)

		j, err := json.Marshal(aims.Aims)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(j))
	} else {
		fmt.Fprint(w, "Bad request!")
	}
}

func main() {
	http.HandleFunc("/", ApiHandler)
	http.ListenAndServe(":3000", nil)
}
