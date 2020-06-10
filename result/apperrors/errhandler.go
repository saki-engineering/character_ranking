package apperrors

// ErrorTypeを返すインターフェース
type typeGetter interface {
	Type() ErrorType
}

// GetType ErrorTypeを持つ場合はそれを返し、無ければUnknownを返す
func GetType(err error) ErrorType {
	for {
		if e, ok := err.(typeGetter); ok {
			return e.Type()
		}
		break
	}
	return Unknown
}

// AppError型のMessageを返すインターフェース
type messageGetter interface {
	Log() string
}

// GetMessage AppError型でエラーメッセージをもつ場合はそれを返し、なければ空文字列を返す
func GetMessage(err error) string {
	for {
		if e, ok := err.(messageGetter); ok {
			return e.Log()
		}
		break
	}
	return ""
}
