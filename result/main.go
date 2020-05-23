package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"app/middlewares"

	"github.com/gorilla/mux"
)

func main() {
	port := "7070"
	fmt.Printf("Result Server Listening on port %s\n", port)

	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)

	fs := http.FileServer(http.Dir("./resources"))
	r.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", fs))

	r.Use(middlewares.Logging)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

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
	//fmt.Fprintf(w, "Hello World")
}

func loadTemplate(name string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(
		"templates/"+name+".html",
		"templates/partials/_header.html",
		"templates/partials/_footer.html",
	)
	return tmpl, err
}
