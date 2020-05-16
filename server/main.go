package main

import (
	"fmt"
	"log"
	"net/http"

	"./handlers"
	"./middlewares"
	"github.com/gorilla/mux"
)

func main() {
	port := "8080"
	fmt.Printf("Server Listening on port %s\n", port)

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.ViewTopHandler)
	r.HandleFunc("/about", handlers.ViewAboutHandler)
	r.HandleFunc("/faq", handlers.ViewFaqHandler)

	s1 := r.PathPrefix("/characters").Subrouter()
	s1.HandleFunc("", handlers.CharacterHandler)
	s1.HandleFunc("/{name}", handlers.CharacterDetailHandler)
	s1.HandleFunc("/{name}/vote", handlers.CharacterVoteHandler).Methods("POST")
	s1.HandleFunc("/{name}/voted", handlers.CharacterVotedHandler)

	s2 := r.PathPrefix("/form").Subrouter()
	s2.HandleFunc("", handlers.FormHandler)
	s2.HandleFunc("/vote", handlers.FormVoteHandler).Methods("POST")

	fs := http.FileServer(http.Dir("./resources"))
	r.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", fs))

	r.Use(middlewares.Logging)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
