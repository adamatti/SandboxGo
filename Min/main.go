package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arturovm/min"
)

func main() {
	m := min.New(nil)

	m.Get("/", helloWorld)

	log.Println("App starded")
	http.ListenAndServe(":8080", m)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello world called")
	fmt.Fprintln(w, "hello world!")
}
