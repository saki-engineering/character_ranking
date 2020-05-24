package handlers

import (
	"log"
	"net/http"
)

// ResultRootHandler /resultのハンドラ
func ResultRootHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("result/index")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := Page{"View Result!", "", false}
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on ResultRootHandler: ", err)
	}
}
