package handlers

import (
	"app/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// CreateUserHandler /user/のPOSTハンドラ
// テストするためには $curl -X POST -d "age=1&gender=1&address=1" localhost:9090/user/
func CreateUserHandler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	db, e := models.ConnectDB()
	if e != nil {
		log.Fatal("connect DB: ", e)
	}
	defer db.Close()

	id, err := models.InsertUsers(db, req.Form.Get("age"), req.Form.Get("gender"), req.Form.Get("address"))
	if err != nil {
		log.Println("insert: ", err)
	} else {
		log.Println("user insert success, id:", id)
	}
	printid := strconv.FormatInt(id, 10)
	fmt.Fprintf(w, printid)
}
