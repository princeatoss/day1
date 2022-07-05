package main

import (
	"encoding/json"
	"fmt"
	"httpserver/entity"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

var person entity.Person

func main() {
	router := chi.NewRouter()
	router.Route("/v1", func(r chi.Router) {
		r.Get("/person", GetPerson)
		r.Post("/person", CreatePerson)

	})
	router.Route("/v2", func(r chi.Router) {
		r.Get("/person", GetPerson)
		r.Post("/person", CreatePerson2)

	})
	http.ListenAndServe(":8080", router)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)
	value := fmt.Sprintf("Person is %v", person)
	fmt.Fprint(w, value)
}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	decodedBody, _ := io.ReadAll(body)
	json.Unmarshal(decodedBody, &person)
	value := fmt.Sprintf("My name is %v", person)
	fmt.Fprint(w, value)
}
func CreatePerson2(w http.ResponseWriter, r *http.Request) {
	person = entity.Person{
		Name:   "Andrei",
		ID:     "1",
		Gender: "Male",
	}
	value := fmt.Sprintf("My name is %v", person)
	fmt.Fprint(w, value)
}
