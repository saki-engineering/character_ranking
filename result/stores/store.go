package stores

import (
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
