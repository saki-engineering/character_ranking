package main

import (
	"fmt"
	"log"
	"net/http"

	"app/routers"
)

func main() {
	port := "8080"
	fmt.Printf("Web Server Listening on port %s\n", port)

	r := routers.CreateRouter()

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
