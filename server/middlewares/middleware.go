package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"app/stores"

	"github.com/google/uuid"
)

// Logging アクセス時にリクエスト内容をロギング
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		fmt.Printf("method: %s || path: %s || form: %s\n", req.Method, req.URL.Path, req.Form)

		next.ServeHTTP(w, req)
	})
}

// CheckSessionID セッションIDがなければ付与する
func CheckSessionID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		id, err := stores.GetSessionID(req)

		if err != nil || id == "" {
			uuid, err2 := uuid.NewRandom()
			if err2 != nil {
				log.Println("cannot make uuid: ", err2)
			}
			cookie := &http.Cookie{
				Name:    stores.SessionName,
				Value:   uuid.String(),
				Expires: time.Now().AddDate(1, 0, 0),
			}
			http.SetCookie(w, cookie)
		}

		next.ServeHTTP(w, req)
	})
}
