package apperrors

// ErrorType カスタムエラーごとに付与するエラーコード
type ErrorType uint

const (
	// Unknown 判別不可時のコード
	Unknown ErrorType = 0
	// DBConnectionFailed DB接続に関するエラーコード
	DBConnectionFailed ErrorType = 1
	// MySQLError package mysqlに関連するエラーコード
	MySQLError ErrorType = 2
)

// Wrap 発生したerrに、エラーコードとカスタムメッセージを与えて、AppError型にするメソッド
func (typecode ErrorType) Wrap(err error, message string) error {
	return AppError{Code: typecode, error: err, Message: message}
}
