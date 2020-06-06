package routers

import (
	"app/handlers"
	"app/middlewares"

	"github.com/gorilla/mux"
)

// CreateRouter muxのルーターを作る
func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.ViewTopHandler)
	r.HandleFunc("/about", handlers.ViewAboutHandler)
	r.HandleFunc("/faq", handlers.ViewFaqHandler)

	s1 := r.PathPrefix("/characters").Subrouter()
	s1.HandleFunc("", handlers.CharacterHandler)
	s1.HandleFunc("/{name}", handlers.CharacterDetailHandler)
	s1.HandleFunc("/{name}/vote", handlers.CharacterVoteHandler).Methods("POST")
	s1.HandleFunc("/{name}/voted", handlers.CharacterVotedHandler)

	s2 := r.PathPrefix("/form").Subrouter()
	s2.HandleFunc("", handlers.FormHandler)
	s2.HandleFunc("/vote", handlers.FormVoteHandler).Methods("POST")

	r.Use(middlewares.Logging)

	return r
}
