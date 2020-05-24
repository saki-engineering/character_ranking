package middlewares

import (
	"fmt"
	"net/http"
)

// Logging アクセス時にリクエスト内容をロギング
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		fmt.Printf("method: %s || path: %s || form: %s\n", req.Method, req.URL.Path, req.Form)

		next.ServeHTTP(w, req)
	})
}

// AuthAdmin ユーザーログインしているかどうかをチェック
func AuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Printf("auth\n")

		next.ServeHTTP(w, req)
	})
}
