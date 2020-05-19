package main

import (
	"fmt"
	"log"
	"net/http"

	"./middlewares"
	"./routers"
)

func main() {
	port := "8080"
	fmt.Printf("Web Server Listening on port %s\n", port)

	r := routers.CreateRouter()

	fs := http.FileServer(http.Dir("./resources"))
	r.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", fs))

	r.Use(middlewares.Logging)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
