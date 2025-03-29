package server

import "fmt"

type HTTPError struct {
	Message string
	Err     error
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func (e *HTTPError) Unwrap() error {
	return e.Err
}

type RouterError struct {
}

func (e *RouterError) Error() string {
	return "failed to get router from server"
}
