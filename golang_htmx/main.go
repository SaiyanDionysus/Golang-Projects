package main

import (
	"fmt"
	"log"
	"text/template"

	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}
		tmpl.Execute(w, films)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

//15:35
