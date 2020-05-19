package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	port := "9090"
	fmt.Printf("API Server Listening on port %s\n", port)

	_, e := sql.Open("mysql", "root@/sampledb")
	if e != nil {
		log.Fatal("DB: ", e)
	} else {
		log.Println("Connected to mysql.")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)

	s1 := r.PathPrefix("/vote").Subrouter()
	s1.HandleFunc("/", VoteRootHandler)
	s1.HandleFunc("/{name}", CharaResultHandler).Methods("GET")
	s1.HandleFunc("/{name}", VoteCharaHandler).Methods("POST")

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// RootHandler /のハンドラ
func RootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// VoteRootHandler /vote/のハンドラ
func VoteRootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Vote root")
}

// CharaResultHandler /vote/{name}のGETハンドラ
func CharaResultHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fmt.Fprintf(w, vars["name"])
}

// VoteCharaHandler /vote/{name}のPOSTハンドラ
// テストするためには $curl -X POST -d "character=cinamon" localhost:9090/vote/name
func VoteCharaHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fmt.Fprintf(w, vars["name"])

	req.ParseForm()
	fmt.Println("form: ", req.Form)
}
