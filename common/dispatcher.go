package common

// Dispatcher interface
type Dispatcher interface {
	Dispatch(cmd Command, data interface{}) (interface{}, error)
}

// Controller
type dispatcher struct{}

// NewDispatcher creates a new dispatcher
func NewDispatcher() Dispatcher {
	return &dispatcher{}
}

// Dispatch a command
func (d *dispatcher) Dispatch(cmd Command, context interface{}) (interface{}, error) {
	err := cmd.CanExecute(context)
	if err != nil {
		return nil, err
	}

	// Execution
	return cmd.Execute(context)
}
