package apperrors

// AppError resultアプリで独自に定義したエラー
type AppError struct {
	Err     error
	Code    ErrorType
	Message string
}

// Error AppError型を、errorインターフェースに代入できるようにする
func (err *AppError) Error() string {
	return err.Message
}

func (err *AppError) Unwrap() error {
	return err.Err
}
