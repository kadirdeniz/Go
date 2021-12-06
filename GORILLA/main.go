package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "End Point Called")
}

func requestHandler() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	fmt.Println("Server Started")

	requestHandler()
}
