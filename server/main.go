package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	fmt.Println("form: ", req.Form)
	fmt.Println("path: ", req.URL.Path)
	fmt.Println("scheme: ", req.URL.Scheme)
	fmt.Println(req.Form["url_long"])
	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello server!")
}

type Page struct {
	Title string
	Count int
}

func viewHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
		panic(err)
	}

	page := Page{"Hello World.", 1}
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute: ", err)
	}
}

func main() {
	fmt.Println("Server Start")

	http.HandleFunc("/", viewHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
