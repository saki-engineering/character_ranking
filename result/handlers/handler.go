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
		apperrors.ErrorHandler(err)
	}

	page := new(Page)
	page.Title = "View Result!"
	page.UserID = ""
	page.LogIn = false
	page.Admin = false
	conn, e := stores.ConnectRedis()
	if e != nil {
		apperrors.ErrorHandler(e)
	}
	defer conn.Close()
	sessionID, _ := stores.GetSessionID(req)

	if userid, _ := stores.GetSessionValue(sessionID, "userid", conn); userid != "" {
		page.UserID = userid
		page.LogIn = true
	}

	if auth, _ := stores.GetSessionValue(sessionID, "auth", conn); auth == "true" {
		page.Admin = true
	}

	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(err)
	}
}

// LoginPageHandler /loginのGETハンドラ
func LoginPageHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("login")
	if err != nil {
		apperrors.ErrorHandler(err)
	}

	page := new(Page)
	page.Title = "View Result!"
	err = executeTemplate(w, tmpl, page)
	if err != nil {
		apperrors.ErrorHandler(err)
	}
}

// LoginHandler /loginのPOSTハンドラ
func LoginHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	db, e := models.ConnectDB()
	if e != nil {
		apperrors.ErrorHandler(e)
	}
	defer db.Close()

	user, err := models.GetUserData(db, req.Form.Get("userid"))
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	err2 := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Form.Get("password")))
	// パスワードが正しくなかった場合はerrが返る
	if err2 != nil {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	conn, err3 := stores.ConnectRedis()
	if err3 != nil {
		apperrors.ErrorHandler(err3)
	}
	defer conn.Close()
	oldSessionID, _ := stores.GetSessionID(req)

	newSessionID, _ := stores.SetSessionID(w)
	stores.ReNameSessionID(oldSessionID, newSessionID, conn)

	stores.SetSessionValue(newSessionID, "userid", user.UserID, conn)
	if user.Auth == 1 {
		stores.SetSessionValue(newSessionID, "auth", "true", conn)
	} else {
		stores.SetSessionValue(newSessionID, "auth", "false", conn)
	}

	log.Println("login success: userid=", user.UserID)

	http.Redirect(w, req, "/", http.StatusSeeOther)
}

// LogoutHandler /logoutのハンドラ
func LogoutHandler(w http.ResponseWriter, req *http.Request) {
	conn, err := stores.ConnectRedis()
	if err != nil {
		apperrors.ErrorHandler(err)
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

	db, e := models.ConnectDB()
	if e != nil {
		apperrors.ErrorHandler(e)
	}
	defer db.Close()

	exist, err := models.CheckIDExist(db, req.Form.Get("userid"))
	if err != nil {
		apperrors.ErrorHandler(err)
	}

	var printnum string
	if exist {
		printnum = "1"
	} else {
		printnum = "0"
	}

	fmt.Fprintf(w, printnum)
}
