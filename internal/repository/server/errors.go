package server

import "fmt"

type ServerError struct {
	Message string
	Err     error
}

func (e *ServerError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func (e *ServerError) Unwrap() error {
	return e.Err
}

type RouterError struct {
}

func (e *RouterError) Error() string {
	return "failed to get router from server"
}
