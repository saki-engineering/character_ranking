package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
)

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

// NameHandler /[name]のハンドラ
func NameHandler(w http.ResponseWriter, req *http.Request) {
	printLogs(req)

	vars := mux.Vars(req)
	fmt.Fprintf(w, "gorilla mux %s", vars["name"])
}

// ViewHandler /のハンドラ
func ViewHandler(w http.ResponseWriter, req *http.Request) {
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
