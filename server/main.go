package main

import (
	"fmt"
	"log"
	"net/http"

	"./handlers"
	"github.com/gorilla/mux"
)

func main() {
	port := "8080"
	fmt.Printf("Server Listening on port %s\n", port)

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.ViewTopHandler)
	r.HandleFunc("/about", handlers.ViewAboutHandler)
	r.HandleFunc("/faq", handlers.ViewFaqHandler)

	r.HandleFunc("/characters", handlers.ViewCharacterHandler)
	r.HandleFunc("/characters/{name}", handlers.NameHandler)

	r.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", http.FileServer(http.Dir("./resources"))))

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
