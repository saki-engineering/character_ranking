package handlers

import (
	"log"
	"net/http"
)

// ViewTopHandler /のハンドラ
func ViewTopHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("index")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := new(Page)
	page.Title = "Character rankinig!"
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

	page := new(Page)
	page.Title = "Character rankinig!"
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

	page := new(Page)
	page.Title = "Character rankinig!"
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on viewHandler: ", err)
	}
}
