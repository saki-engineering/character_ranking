package routers

import (
	"app/handlers"

	"github.com/gorilla/mux"
)

// CreateRouter muxのルーターを作る
func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.RootHandler)

	s1 := r.PathPrefix("/vote").Subrouter()
	s1.HandleFunc("/", handlers.VoteResultHandler).Methods("GET")
	s1.HandleFunc("/", handlers.VoteCharaHandler).Methods("POST")
	s1.HandleFunc("/summary", handlers.VoteSammaryHandler).Methods("GET")
	s1.HandleFunc("/{name}", handlers.CharaResultHandler).Methods("GET")

	s2 := r.PathPrefix("/user").Subrouter()
	s2.HandleFunc("/", handlers.CreateUserHandler).Methods("POST")

	return r
}
