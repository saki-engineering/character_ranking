package routers

import (
	"app/handlers"
	"app/middlewares"

	"github.com/gorilla/mux"
)

// CreateRouter muxのルーターを作る
func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.RootHandler)
	r.HandleFunc("/login", handlers.LoginPageHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler)

	s1 := r.PathPrefix("/result").Subrouter()
	s1.HandleFunc("", handlers.ResultRootHandler)

	s2 := r.PathPrefix("/admin").Subrouter()
	s2.HandleFunc("", handlers.AdminRootHandler)
	s2.HandleFunc("/userform", handlers.CreateUserFormHandler).Methods("GET")
	s2.HandleFunc("/userform", handlers.CreateUserHandler).Methods("POST")

	r.Use(middlewares.Logging)
	s1.Use(middlewares.AuthAdmin)
	s2.Use(middlewares.AuthSuperAdmin)

	return r
}
