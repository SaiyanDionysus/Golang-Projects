package main

import (
	"fmt"
	"io"
	"log"

	"net/http"
)

func main() {
	fmt.Println("Hello")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world\n")
		io.WriteString(w, r.Method)
	}

	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
