package main

import (
	"fmt"
	"log"
	"net/http"

	"app/middlewares"

	"github.com/gorilla/mux"
)

func main() {
	port := "7070"
	fmt.Printf("Result Server Listening on port %s\n", port)

	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)

	r.Use(middlewares.Logging)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// RootHandler /のハンドラ
func RootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
