package db

import "fmt"

type DBError struct {
	Message string
	Err     error
}

func (e *DBError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func (e *DBError) Unwrap() error {
	return e.Err
}
