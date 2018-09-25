package errors

type errorType string

const (
	// UnknownError type
	UnknownError errorType = "UnknownError"
	// ValidationError type
	ValidationError errorType = "ValidationError"
	// AuthenticationError type
	AuthenticationError errorType = "AuthenticationError"
	// AuthorizationError type
	AuthorizationError errorType = "AuthorizationError"
)

// CustomError struct
type CustomError struct {
	Kind errorType
	Msg  string
}

// New create a custom error
func New(kind errorType, msg string) *CustomError {
	return &CustomError{Kind: kind, Msg: msg}
}

// Error returns error message
func (err *CustomError) Error() string {
	return string(err.Kind) + ": " + err.Msg
}

// // Log returns a string with the error type and the error message
// func (err *CustomError) Log() string {
// 	return string(err.Kind) + ": " + err.Msg
// }
