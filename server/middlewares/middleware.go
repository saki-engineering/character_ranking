package middlewares

import (
	"fmt"
	"net/http"

	"app/stores"
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
			stores.SetSessionID(w)
		}

		next.ServeHTTP(w, req)
	})
}
