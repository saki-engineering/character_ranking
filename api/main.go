package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := "9090"
	fmt.Printf("API Server Listening on port %s\n", port)

	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// RootHandler /のハンドラ
func RootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
