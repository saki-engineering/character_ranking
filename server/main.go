package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	port := "8080"
	fmt.Printf("Server Listening on port %s\n", port)

	http.HandleFunc("/", viewHandler)
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
	Title string
	Count int
}

func viewHandler(w http.ResponseWriter, req *http.Request) {
	printLogs(req)

	tmpl, err := loadTemplate("layout")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := Page{"Character rankinig!", 1}
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
