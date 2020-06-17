package handlers

import (
	"app/apperrors"

	"net/http"
)

// ViewTopHandler /のハンドラ
func ViewTopHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("index")
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	page := new(Page)
	page.Title = "Character rankinig!"
	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
}

// ViewAboutHandler Aboutページのハンドラ
func ViewAboutHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("about")
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	page := new(Page)
	page.Title = "Character rankinig!"
	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
}

// ViewFaqHandler FAQページのハンドラ
func ViewFaqHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("faq")
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	page := new(Page)
	page.Title = "Character rankinig!"
	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
}
