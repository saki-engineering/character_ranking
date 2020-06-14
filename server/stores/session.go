package stores

import (
	"net/http"
	"time"

	"app/apperrors"

	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

var (
	// SessionName 情報を保持するキー
	SessionName = "votesessionID"
)

// ConnectRedis redisと接続する
func ConnectRedis() (redis.Conn, error) {
	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		apperrors.SessionStrageConnectionFailed.Wrap(err, "cannot use session")
		return nil, err
	}
	return conn, nil
}

// GetSessionID セッションIDを取得
func GetSessionID(req *http.Request) (string, error) {
	cookie, err := req.Cookie(SessionName)
	if err != nil {
		err = apperrors.SessionIDGetFailed.Wrap(err, "cannot use session")
		return "", err
	}
	return cookie.Value, nil
}

// SetSessionID セッションIDを生成して、cookieにセットする
// 返り値は生成したセッションID
func SetSessionID(w http.ResponseWriter) (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		err = apperrors.SessionIDCreatedFailed.Wrap(err, "cannot use session")
		return "", err
	}

	newSessionID := uuid.String()
	cookie := &http.Cookie{
		Name:    SessionName,
		Value:   newSessionID,
		Expires: time.Now().AddDate(1, 0, 0),
	}
	http.SetCookie(w, cookie)

	return newSessionID, nil
}

// GetSessionValue セッションIDとfieldからvalueを取得
func GetSessionValue(sessionID, field string, conn redis.Conn) (string, error) {
	value, err := redis.String(conn.Do("HGET", sessionID, field))
	if err != nil {
		err = apperrors.SessionInfoEditFailed.Wrap(err, "failed to edit session data")
	}
	return value, err
}

// SetSessionValue セッションIDを受け取って、(field,value)の組をセット
func SetSessionValue(sessionID, field, value string, conn redis.Conn) error {
	ttl := 86400

	_, err := conn.Do("HSET", sessionID, field, value)
	if err != nil {
		err = apperrors.SessionInfoEditFailed.Wrap(err, "failed to edit session data")
	}
	conn.Do("EXPIRE", sessionID, ttl)
	return err
}

// DeleteSessionValue セッションIDを受け取って、fieldを削除
func DeleteSessionValue(sessionID, field string, conn redis.Conn) error {
	_, err := conn.Do("HDEL", sessionID, field)
	if err != nil {
		err = apperrors.SessionInfoEditFailed.Wrap(err, "failed to edit session data")
	}
	return err
}
