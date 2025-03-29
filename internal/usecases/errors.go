package usecases

type ServiceError struct {
	Message string
}

func (e *ServiceError) Error() string {
	return e.Message
}
