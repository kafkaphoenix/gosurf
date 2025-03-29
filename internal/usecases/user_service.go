package usecases

import (
	"fmt"

	"github.com/kafkaphoenix/gosurf/internal/domain"
	"github.com/kafkaphoenix/gosurf/internal/repository"
)

type UserService struct {
	db *repository.FakeDB
}

func NewUserService(db *repository.FakeDB) *UserService {
	return &UserService{db: db}
}

// GetUserByID return a user if exists.
func (s *UserService) GetUserByID(userID int) (*domain.User, error) {
	user, exists := s.db.Users[userID]
	if !exists {
		return nil, &ServiceError{Message: fmt.Sprintf("user with id %d not found", userID)}
	}

	return &user, nil
}
