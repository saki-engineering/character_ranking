package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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
