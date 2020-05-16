package handlers

import (
	"log"
	"net/http"
	"text/template"
)

//Page ... htmlに渡す値をまとめた構造体
type Page struct {
	Title     string
	Character []string
}

var charas = []string{
	"cinnamon",
	"cappuccino",
	"mocha",
	"chiffon",
	"espresso",
	"milk",
	"azuki",
	"coco",
	"nuts",
	"poron",
	"corne",
	"berry",
	"cherry",
}

// ViewTopHandler /のハンドラ
func ViewTopHandler(w http.ResponseWriter, req *http.Request) {
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

// ViewAboutHandler Aboutページのハンドラ
func ViewAboutHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("about")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := Page{"About", charas}
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on viewHandler: ", err)
	}
}

// ViewFaqHandler FAQページのハンドラ
func ViewFaqHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("faq")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := Page{"FAQ", charas}
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on viewHandler: ", err)
	}
}

func loadTemplate(name string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(
		"templates/"+name+".html",
		"templates/partials/_header.html",
		"templates/partials/_footer.html",
	)
	return tmpl, err
}
