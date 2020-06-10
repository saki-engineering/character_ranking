package apperrors

// ErrorType カスタムエラーごとに付与するエラーコード
type ErrorType uint

const (
	// Unknown 判別不可時のコード
	Unknown ErrorType = 0
	// DBConnectionFailed DB接続に関するエラーコード
	DBConnectionFailed ErrorType = 1
	// MySQLSetUpError package mysqlのテーブル作成に関連するエラーコード
	MySQLSetUpError ErrorType = 21
	// MySQLExecError SQL文実行時に何かあったときのエラーコード
	MySQLExecError ErrorType = 22
	// MySQLQueryError SQL文でデータを取得できなかったときのエラーコード
	MySQLQueryError ErrorType = 23
	// MySQLDataCreateFailed 挿入するデータをアプリ側で作成するのに失敗したときのエラーコード
	MySQLDataCreateFailed ErrorType = 24
	// MySQLDataFormatFailed DBから取得したデータのScan→golangの構造体に直すのに失敗した時のエラーコード
	MySQLDataFormatFailed ErrorType = 25
)

// Wrap 発生したerrに、エラーコードとカスタムメッセージを与えて、AppError型にするメソッド
func (typecode ErrorType) Wrap(err error, message string) error {
	return AppError{Code: typecode, error: err, Message: message}
}
