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
	// MySQLExecError SQL文実行時に何かあったときのエラーコード
	MySQLExecError ErrorType = 21
	// MySQLQueryError SQL文でデータを取得できなかったときのエラーコード
	MySQLQueryError ErrorType = 22
	// MySQLDataCreateFailed 挿入するデータをアプリ側で作成するのに失敗したときのエラーコード
	MySQLDataCreateFailed ErrorType = 23
	// MySQLDataFormatFailed DBから取得したデータのScan→golangの構造体に直すのに失敗した時のエラーコード
	MySQLDataFormatFailed ErrorType = 24

	// SessionStrageConnectionFailed セッション情報を保存するredis接続に関するエラーコード
	SessionStrageConnectionFailed ErrorType = 30
	// SessionIDGetFailed ユーザーのcookieからIDが取得できなかったときのエラーコード
	SessionIDGetFailed ErrorType = 31
	// SessionIDCreatedFailed uuidでのランダムセッションID生成に失敗したときのエラーコード
	SessionIDCreatedFailed ErrorType = 32
	// SessionInfoEditFailed redisコマンド実行時のエラーコード
	SessionInfoEditFailed ErrorType = 33

	// HTMLTemplateLoadFailed テンプレートの読み込みに失敗したときのエラーコード
	HTMLTemplateLoadFailed ErrorType = 40
	// HTMLTemplateExecFailed テンプレートにPage構造体を渡すのに失敗したときのエラーコード
	HTMLTemplateExecFailed ErrorType = 41

	// VoteAPIRequestError vote_apiのリクエストに失敗したときのエラーコード
	VoteAPIRequestError ErrorType = 50
	// VoteAPIResponseReadFailed vote_apiからのレスポンスボディの読み込み失敗orjsonパース失敗時のエラーコード
	VoteAPIResponseReadFailed ErrorType = 51

	// HTTPServerPortListenFailed httpサーバーの軌道に失敗したときのエラーコード
	HTTPServerPortListenFailed ErrorType = 60
)

// Wrap 発生したerrに、エラーコードとカスタムメッセージを与えて、AppError型にするメソッド
func (typecode ErrorType) Wrap(err error, message string) error {
	return AppError{Code: typecode, error: err, Message: message}
}
