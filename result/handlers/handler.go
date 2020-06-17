package handlers

import (
	"fmt"
	"log"
	"net/http"

	"app/apperrors"
	"app/models"
	"app/stores"

	"golang.org/x/crypto/bcrypt"
)

// RootHandler /のハンドラ
func RootHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("index")
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	page := new(Page)
	page.Title = "View Result!"
	page.UserID = ""
	page.LogIn = false
	page.Admin = false
	conn, err := stores.ConnectRedis()
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	defer conn.Close()
	sessionID, _ := stores.GetSessionID(req)

	if nowLoginUserID, _ := stores.GetSessionValue(sessionID, "userid", conn); nowLoginUserID != "" {
		page.UserID = nowLoginUserID
		page.LogIn = true
	}

	if nowLoginUserAuth, _ := stores.GetSessionValue(sessionID, "auth", conn); nowLoginUserAuth == "true" {
		page.Admin = true
	}

	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
}

// LoginPageHandler /loginのGETハンドラ
func LoginPageHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("login")
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	page := new(Page)
	page.Title = "View Result!"
	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
}

// LoginHandler /loginのPOSTハンドラ
func LoginHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	formInputUserID := req.Form.Get("userid")
	formInputUserPlainPassword := req.Form.Get("password")

	db, err := models.ConnectDB()
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	defer db.Close()

	loginTryingUser, err := models.GetUserData(db, formInputUserID)
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(loginTryingUser.HashedPassword), []byte(formInputUserPlainPassword))
	// パスワードが正しくなかった場合はerrが返る
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	conn, err := stores.ConnectRedis()
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	defer conn.Close()
	oldSessionID, _ := stores.GetSessionID(req)

	newSessionID, _ := stores.SetSessionID(w)
	stores.ReNameSessionID(oldSessionID, newSessionID, conn)

	stores.SetSessionValue(newSessionID, "userid", loginTryingUser.UserID, conn)
	if loginTryingUser.Auth == 1 {
		stores.SetSessionValue(newSessionID, "auth", "true", conn)
	} else {
		stores.SetSessionValue(newSessionID, "auth", "false", conn)
	}

	log.Println("login success: userid=", loginTryingUser.UserID)

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

// LogoutHandler /logoutのハンドラ
func LogoutHandler(w http.ResponseWriter, req *http.Request) {
	conn, err := stores.ConnectRedis()
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	defer conn.Close()
	sessionID, _ := stores.GetSessionID(req)

	stores.DeleteSessionValue(sessionID, "userid", conn)
	stores.DeleteSessionValue(sessionID, "auth", conn)

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

// CheckIDHandler /checkidのPOSTハンドラ
func CheckIDHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	formInputUserID := req.Form.Get("userid")

	db, err := models.ConnectDB()
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	defer db.Close()

	exist, err := models.CheckIDExist(db, formInputUserID)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	var printnum string
	if exist {
		printnum = "1"
	} else {
		printnum = "0"
	}

	fmt.Fprintf(w, printnum)
}
