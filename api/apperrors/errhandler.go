package apperrors

import (
	"errors"
	"log"
	"net/http"
)

// ErrorHandler エラーが発生したときのロギング・アプリの終了判定をここで一括で行う
func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		log.Printf("[AppError] ErrorType %d: %s", appErr.Code, errors.Unwrap(appErr))
	} else {
		appErr = &AppError{Err: err, Code: Unknown, Message: "Unknown Error occured"}
		log.Printf("[AppError] ErrorType %d: %s", appErr.Code, errors.Unwrap(appErr))
	}
	http.Error(w, appErr.Error(), http.StatusInternalServerError)
}
