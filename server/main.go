package main

import (
	"fmt"
	"log"
	"net/http"

	"app/routers"
	"app/stores"
)

func main() {
	port := "8080"
	fmt.Printf("Web Server Listening on port %s\n", port)

	r := routers.CreateRouter()

	fs := http.FileServer(http.Dir("./resources"))
	r.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", fs))

	conn, e := stores.ConnectRedis()
	if e != nil {
		log.Fatal("cannot connect redis: ", e)
	} else {
		log.Println("suuccess to connect redis")
	}
	defer conn.Close()

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
