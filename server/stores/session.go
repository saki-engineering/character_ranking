package stores

import (
	"log"
	"net/http"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

var (
	// SessionName 情報を保持するキー
	SessionName = "votesessionID"
)

// ConnectRedis redisと接続する
func ConnectRedis() (redis.Conn, error) {
	return redis.Dial("tcp", "redis:6379")
}

// GetSessionID セッションIDを取得
func GetSessionID(req *http.Request) (string, error) {
	cookie, err := req.Cookie(SessionName)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

// SetSessionID セッションIDを生成して、cookieにセットする
func SetSessionID(w http.ResponseWriter) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Println("cannot make uuid: ", err)
	}

	cookie := &http.Cookie{
		Name:    stores.SessionName,
		Value:   uuid.String(),
		Expires: time.Now().AddDate(1, 0, 0),
	}
	http.SetCookie(w, cookie)
}

// GetSessionValue セッションIDとfieldからvalueを取得
func GetSessionValue(sessionID, field string, conn redis.Conn) (string, error) {
	value, err := redis.String(conn.Do("HGET", sessionID, field))
	return value, err
}

// SetSessionValue セッションIDを受け取って、(field,value)の組をセット
func SetSessionValue(sessionID, field, value string, conn redis.Conn) error {
	ttl := 86400

	_, err := conn.Do("HSET", sessionID, field, value)
	conn.Do("EXPIRE", sessionID, ttl)
	return err
}

// DeleteSessionValue セッションIDを受け取って、fieldを削除
func DeleteSessionValue(sessionID, field string, conn redis.Conn) error {
	_, err := conn.Do("HDEL", sessionID, field)
	return err
}
