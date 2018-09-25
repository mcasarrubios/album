package common

// Command interface
type Command interface {
	Execute(context interface{}) (interface{}, error)
	CanExecute(context interface{}) error
}
