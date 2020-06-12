package apperrors

// AppError resultアプリで独自に定義したエラー
type AppError struct {
	error   // フィールド名を省略→型名と同じフィールド名に自動的になる
	Code    ErrorType
	Message string
}

// Error AppError型を、errorインターフェースに代入できるようにする
func (err AppError) Error() string {
	return err.error.Error()
}

// Type AppError型から、エラーコードを入手するメソッド
// → typeGetterインターフェースに代入可能に
func (err AppError) Type() ErrorType {
	return err.Code
}

// Log AppError型から、エラーメッセージを入手するメソッド
// → messageGetterインターフェースに代入可能に
func (err AppError) Log() string {
	return err.Message
}
