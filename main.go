package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty`
	State string `json:state,omitempty`
}

var people []Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	id := len(people) + 1
	person.ID = strconv.Itoa(id)
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
	data()
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	parmas := mux.Vars(r)

	for index, item := range people {
		if item.ID == parmas["id"] {
			people = append(people[:index], people[index+1:]...)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func data() {
	fmt.Printf("Data: %v\n", people)
}

func main() {

	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Z"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
	people = append(people, Person{ID: "4", Firstname: "Luke"})

	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	http.Handle("/", router)

	fmt.Println("HTTP Server started on :8000")

	data()

	log.Fatal(http.ListenAndServe(":8000", router))
}
