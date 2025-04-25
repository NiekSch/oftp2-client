package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"oftp2-client/cmd/oftp2/API"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/users", API.GetUsers).Methods("GET")
	// router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	fmt.Fprintf(w, "ID is: %s", vars["id"])
	// }).Methods("GET")
	router.HandleFunc("/users/{id}", API.GetUser).Methods("GET")
	router.HandleFunc("/determineId/{SSID}", API.DetermineId).Methods("GET")
	router.HandleFunc("/sendFile", API.SendFile).Methods("GET")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
