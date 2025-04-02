package db

import "fmt"

type FakeDBError struct {
	Message string
	Err     error
}

func (e *FakeDBError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func (e *FakeDBError) Unwrap() error {
	return e.Err
}
