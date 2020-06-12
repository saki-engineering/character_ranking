package handlers

import (
	"app/apperrors"

	"net/http"
)

// ViewTopHandler /のハンドラ
func ViewTopHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("index")
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	page := new(Page)
	page.Title = "Character rankinig!"
	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
}

// ViewAboutHandler Aboutページのハンドラ
func ViewAboutHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("about")
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	page := new(Page)
	page.Title = "Character rankinig!"
	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
}

// ViewFaqHandler FAQページのハンドラ
func ViewFaqHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("faq")
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	page := new(Page)
	page.Title = "Character rankinig!"
	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
}
