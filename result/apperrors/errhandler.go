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
