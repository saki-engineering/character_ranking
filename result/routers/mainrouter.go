package routers

import (
	"app/handlers"
	"app/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateRouter muxのルーターを作る
func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.RootHandler)
	r.HandleFunc("/login", handlers.LoginPageHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler)
	r.HandleFunc("/checkid", handlers.CheckIDHandler).Methods("POST")

	s1 := r.PathPrefix("/result").Subrouter()
	s1.HandleFunc("", handlers.ResultRootHandler)
	s1.HandleFunc("/user", handlers.UserSummaryHandler)
	s1.HandleFunc("/{name}", handlers.CharacterResultHandler)
	s1.HandleFunc("/user/{gender}/{agemin}", handlers.UserDetailHandler)

	s2 := r.PathPrefix("/admin").Subrouter()
	s2.HandleFunc("", handlers.AdminRootHandler)
	s2.HandleFunc("/userform", handlers.CreateUserFormHandler).Methods("GET")
	s2.HandleFunc("/userform", handlers.CreateUserHandler).Methods("POST")
	s2.HandleFunc("/newuser", handlers.CheckUserHandler)

	fs := http.FileServer(http.Dir("./resources"))
	r.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", fs))

	r.Use(middlewares.Logging)
	r.Use(middlewares.CheckSessionID)
	s1.Use(middlewares.AuthAdmin)
	s2.Use(middlewares.AuthSuperAdmin)

	return r
}
