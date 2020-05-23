package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

//Page ... htmlに渡す値をまとめた構造体
type Page struct {
	Title string
}

// RootHandler /のハンドラ
func RootHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("index")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := Page{"View Result!"}
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on RootHandler: ", err)
	}
}

// LoginPageHandler /loginのGETハンドラ
func LoginPageHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("login")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := Page{"View Result!"}
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on RootHandler: ", err)
	}
}

// LoginHandler /loginのPOSTハンドラ
func LoginHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Login POST")
}

// SignupPageHandler /signupのGETハンドラ
func SignupPageHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("signup")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := Page{"View Result!"}
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on RootHandler: ", err)
	}
}

// SignupHandler /signupのPOSTハンドラ
func SignupHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Signup POST")
}

func loadTemplate(name string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(
		"templates/"+name+".html",
		"templates/partials/_header.html",
		"templates/partials/_footer.html",
	)
	return tmpl, err
}
