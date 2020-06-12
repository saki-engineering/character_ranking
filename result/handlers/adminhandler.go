package handlers

import (
	"app/apperrors"
	"app/models"
	"app/stores"

	"log"
	"net/http"
)

// AdminRootHandler /adminのハンドラ
func AdminRootHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("admin/index")
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	page := new(Page)
	page.Title = "View Result!"

	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
}

// CreateUserFormHandler /admin/userformのGETハンドラ
func CreateUserFormHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("admin/form")
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	page := new(Page)
	page.Title = "View Result!"
	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
}

// CreateUserHandler /admin/userformのPOSTハンドラ
func CreateUserHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	conn, err := stores.ConnectRedis()
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	sessionID, _ := stores.GetSessionID(req)

	stores.SetSessionValue(sessionID, "newuserid", req.Form.Get("userid"), conn)
	stores.SetSessionValue(sessionID, "newpassword", req.Form.Get("password"), conn)

	db, err := models.ConnectDB()
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	auth := 0
	if req.Form.Get("admin") == "on" {
		auth = 1
	}

	err = models.UserCreate(db, req.Form.Get("userid"), req.Form.Get("password"), auth)
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
	log.Println("success to create admin user")
	http.Redirect(w, req, "/admin/newuser", http.StatusSeeOther)
}

// CheckUserHandler /admin/newuserのハンドラ
func CheckUserHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("admin/newuser")
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	conn, err := stores.ConnectRedis()
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	sessionID, _ := stores.GetSessionID(req)

	page := new(Page)
	page.Title = "View Result!"
	page.NewUser.UserID, _ = stores.GetSessionValue(sessionID, "newuserid", conn)
	page.NewUser.Password, _ = stores.GetSessionValue(sessionID, "newpassword", conn)

	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(err)
		http.Error(w, apperrors.GetMessage(err), http.StatusInternalServerError)
		return
	}

	stores.DeleteSessionValue(sessionID, "newuserid", conn)
	stores.DeleteSessionValue(sessionID, "newpassword", conn)
}
