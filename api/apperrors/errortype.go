package apperrors

// ErrorType カスタムエラーごとに付与するエラーコード
type ErrorType uint

const (
	// Unknown 判別不可時のコード
	Unknown ErrorType = 0

	// DBConnectionFailed DB接続に関するエラーコード
	DBConnectionFailed ErrorType = 10

	// MySQLSetUpError package mysqlのテーブル作成に関連するエラーコード
	MySQLSetUpError ErrorType = 20
	// MySQLExecError SQL文実行時に何かあったときのエラーコード(LastinsertedIDが取得できないときもこれ)
	MySQLExecError ErrorType = 21
	// MySQLQueryError SQL文でデータを取得できなかったときのエラーコード
	MySQLQueryError ErrorType = 22
	// MySQLDataCreateFailed 挿入するデータをアプリ側で作成するのに失敗したときのエラーコード
	//MySQLDataCreateFailed ErrorType = 23
	// MySQLDataFormatFailed DBから取得したデータのScan→golangの構造体に直すのに失敗した時のエラーコード
	MySQLDataFormatFailed ErrorType = 24

	// JSONFormatFailed goの構造体をjsonに変換するのに失敗したときのエラーコード
	JSONFormatFailed ErrorType = 30

	// HTTPServerPortListenFailed httpサーバーの軌道に失敗したときのエラーコード
	HTTPServerPortListenFailed ErrorType = 40
)

// Wrap 発生したerrに、エラーコードとカスタムメッセージを与えて、AppError型にするメソッド
func (typecode ErrorType) Wrap(err error, message string) error {
	return &AppError{Code: typecode, Err: err, Message: message}
}
