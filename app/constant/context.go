package constant

type ContextKey string

const (
	REQUEST_ID ContextKey = "RequestId"
)

var contextKeysForLog = []ContextKey{REQUEST_ID}

func ContextKeysForLog() []ContextKey {
	return contextKeysForLog
}
