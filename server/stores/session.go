package stores

import (
	"net/http"

	"github.com/gomodule/redigo/redis"
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

// GetSessionValue セッションIDとfieldからvalueを取得
func GetSessionValue(sessionID, field string, conn redis.Conn) (string, error) {
	value, err := redis.String(conn.Do("HGET", sessionID, field))
	return value, err
}

// SetSessionValue セッションIDを受け取って、(field,value)の組をセット
func SetSessionValue(sessionID, field, value string, conn redis.Conn) error {
	_, err := conn.Do("HSET", sessionID, field, value)
	return err
}

// DeleteSessionValue セッションIDを受け取って、fieldを削除
func DeleteSessionValue(sessionID, field string, conn redis.Conn) error {
	_, err := conn.Do("HDEL", sessionID, field)
	return err
}
