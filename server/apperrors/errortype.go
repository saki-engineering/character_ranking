package apperrors

// ErrorType カスタムエラーごとに付与するエラーコード
type ErrorType uint

const (
	// Unknown 判別不可時のコード
	Unknown ErrorType = 0

	// SessionStrageConnectionFailed セッション情報を保存するredis接続に関するエラーコード
	SessionStrageConnectionFailed ErrorType = 10
	// SessionIDGetFailed ユーザーのcookieからIDが取得できなかったときのエラーコード
	SessionIDGetFailed ErrorType = 11
	// SessionIDCreatedFailed uuidでのランダムセッションID生成に失敗したときのエラーコード
	SessionIDCreatedFailed ErrorType = 12
	// SessionInfoEditFailed redisコマンド実行時のエラーコード
	SessionInfoEditFailed ErrorType = 13

	// HTMLTemplateLoadFailed テンプレートの読み込みに失敗したときのエラーコード
	HTMLTemplateLoadFailed ErrorType = 20
	// HTMLTemplateExecFailed テンプレートにPage構造体を渡すのに失敗したときのエラーコード
	HTMLTemplateExecFailed ErrorType = 21

	// VoteAPIRequestError vote_apiのリクエストに失敗したときのエラーコード
	VoteAPIRequestError ErrorType = 30

	// HTTPServerPortListenFailed httpサーバーの軌道に失敗したときのエラーコード
	HTTPServerPortListenFailed ErrorType = 40
)

// Wrap 発生したerrに、エラーコードとカスタムメッセージを与えて、AppError型にするメソッド
func (typecode ErrorType) Wrap(err error, message string) error {
	return &AppError{Code: typecode, Err: err, Message: message}
}
