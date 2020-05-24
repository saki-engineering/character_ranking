package stores

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var (
	// SessionStore セッションの値をサーバー内で保持するデータストア
	SessionStore *sessions.CookieStore

	// SessionName ログイン情報を保持するキー
	SessionName = "loginsession"
)

// SessionInit セッションストアを初期化する
func SessionInit() {
	SessionStore = sessions.NewCookieStore([]byte(os.Getenv("KARI_SESSION_KEY")))
}

// GetSession リクエストからセッションの取得
func GetSession(req *http.Request) (*sessions.Session, error) {
	return SessionStore.Get(req, SessionName)
}
