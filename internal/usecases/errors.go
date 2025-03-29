package usecases

type UserServiceError struct {
	Message string
}

func (e *UserServiceError) Error() string {
	return e.Message
}
