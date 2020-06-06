package handlers

import (
	"app/models"
	"app/stores"
	"log"
	"net/http"
)

// AdminRootHandler /adminのハンドラ
func AdminRootHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("admin/index")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := new(Page)
	page.Title = "View Result!"
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on RootHandler: ", err)
	}
}

// CreateUserFormHandler /admin/userformのGETハンドラ
func CreateUserFormHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("admin/form")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	page := new(Page)
	page.Title = "View Result!"
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on RootHandler: ", err)
	}
}

// CreateUserHandler /admin/userformのPOSTハンドラ
func CreateUserHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	conn, err2 := stores.ConnectRedis()
	if err2 != nil {
		log.Fatal("cannot connect redis: ", err2)
	}
	defer conn.Close()
	sessionID, _ := stores.GetSessionID(req)

	stores.SetSessionValue(sessionID, "newuserid", req.Form.Get("userid"), conn)
	stores.SetSessionValue(sessionID, "newpassword", req.Form.Get("password"), conn)

	db, e := models.ConnectDB()
	if e != nil {
		log.Fatal("connect DB: ", e)
	}
	defer db.Close()

	auth := 0
	if req.Form.Get("admin") == "on" {
		auth = 1
	}

	err := models.UserCreate(db, req.Form.Get("userid"), req.Form.Get("password"), auth)
	if err != nil {
		log.Println("create admin user: ", err)
	} else {
		log.Println("success to create admin user")
	}
	http.Redirect(w, req, "/admin/newuser", http.StatusSeeOther)
}

// CheckUserHandler /admin/newuserのハンドラ
func CheckUserHandler(w http.ResponseWriter, req *http.Request) {
	tmpl, err := loadTemplate("admin/newuser")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	conn, err2 := stores.ConnectRedis()
	if err2 != nil {
		log.Fatal("cannot connect redis: ", err2)
	}
	defer conn.Close()
	sessionID, _ := stores.GetSessionID(req)

	page := new(Page)
	page.Title = "View Result!"
	page.NewUser.UserID, _ = stores.GetSessionValue(sessionID, "newuserid", conn)
	page.NewUser.Password, _ = stores.GetSessionValue(sessionID, "newpassword", conn)

	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal("Execute on RootHandler: ", err)
	}

	stores.DeleteSessionValue(sessionID, "newuserid", conn)
	stores.DeleteSessionValue(sessionID, "newpassword", conn)
}
