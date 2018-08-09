package errors

// HTTP contains error and statusCode
type HTTP struct {
	error
	StatusCode int
}

// type customError {
// 	error
// 	HTTP
// 	operational
// }

// type ErrorHandler interface {
// 	Error() string
// 	isOperational() bool
// 	StatusCode() int
// 	StatusMessage() string
// }

// func New() ErrorHandler {
// 	return
// }

// func (err customError) Error() string {
// 	return http.err.Error()
// }

// func (err customError) isOperational() bool {
// 	return err.operational
// }

// func (err customError) StatusCode() bool {
// 	return http.operational
// }

// func (err customError) isOperational() bool {
// 	return http.operational
// }
