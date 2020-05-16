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
	r.HandleFunc("/form", handlers.ViewFormHandler)
	r.HandleFunc("/form/vote", handlers.FormVoteHandler)

	s := r.PathPrefix("/characters").Subrouter()
	s.HandleFunc("", handlers.ViewCharacterHandler)
	s.HandleFunc("/{name}", handlers.CharacterDetailHandler)
	s.HandleFunc("/{name}/vote", handlers.VoteHandler).Methods("POST")
	s.HandleFunc("/{name}/voted", handlers.VotedHandler)

	fs := http.FileServer(http.Dir("./resources"))
	r.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", fs))

	r.Use(middlewares.Logging)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
