package routers

import (
	"app/handlers"

	"github.com/gorilla/mux"
)

// CreateRouter muxのルーターを作る
func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.RootHandler)
	r.HandleFunc("/login", handlers.LoginPageHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler)
	r.HandleFunc("/signup", handlers.SignupPageHandler).Methods("GET")
	r.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")

	return r
}
