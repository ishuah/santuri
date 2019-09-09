package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Println("Running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("{ 'message': 'Hello!'}"))
}