package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/NiekSch/oftp2-client/cmd/oftp2/API/idController"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}