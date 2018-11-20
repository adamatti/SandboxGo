package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func getPeople(w http.ResponseWriter, r *http.Request) {
	log.Println("Get people called")
	json.NewEncoder(w).Encode(people)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println("Get person called", params)
	for _, item := range people {
		if item.ID == params["id"] {
			log.Println("Found person!")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// json.NewEncoder(w).Encode(nil)
	http.NotFound(w, r)
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println("Create Person called", params)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println("Delete Person called", params)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			//log.Println("Shall remove index", index)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func buildRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/people", getPeople).Methods("GET")
	router.HandleFunc("/people/{id}", getPerson).Methods("GET")
	router.HandleFunc("/people/{id}", createPerson).Methods("POST")
	router.HandleFunc("/people/{id}", deletePerson).Methods("DELETE")

	return router
}

func main() {
	log.Println("App starting")

	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})

	log.Fatal(http.ListenAndServe(":8000", buildRouter()))
}
