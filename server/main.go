package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	port := "8080"
	fmt.Printf("Server Listening on port %s\n", port)

	r := mux.NewRouter()
	r.HandleFunc("/", viewHandler)
	r.HandleFunc("/{name}", nameHandler)
	http.Handle("/", r)

	//http.HandleFunc("/", viewHandler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func printLogs(req *http.Request) {
	req.ParseForm()

	fmt.Println("form: ", req.Form)
	fmt.Println("path: ", req.URL.Path)
	fmt.Println("scheme: ", req.URL.Scheme)
	fmt.Println(req.Form["url_long"])
	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
}

//Page ... htmlに渡す値をまとめた構造体
type Page struct {
	Title     string
	Character []string
}

var charas = []string{
	"Cinnamon",
	"Cappuccino",
	"Mocha",
	"Chiffon",
	"Espresso",
	"Milk",
	"Azuki",
	"Coco",
	"Nuts",
	"Poron",
	"Corne",
	"Berry",
	"Cherry",
}

func nameHandler(w http.ResponseWriter, req *http.Request) {
	printLogs(req)

	vars := mux.Vars(req)
	fmt.Fprintf(w, "gorilla mux %s", vars["name"])
}

func viewHandler(w http.ResponseWriter, req *http.Request) {
	printLogs(req)

	tmpl, err := loadTemplate("index")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := Page{"Character rankinig!", charas}
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on viewHandler: ", err)
	}
}

func loadTemplate(name string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(
		"templates/"+name+".html",
		"templates/_header.html",
		"templates/_footer.html",
	)
	return tmpl, err
}
