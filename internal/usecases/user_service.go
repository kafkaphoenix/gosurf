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

// GetTotalActionsByID return total actions of a user if exists.
func (s *UserService) GetTotalActionsByID(userID int) (*domain.TotalActions, error) {
	_, exists := s.db.Users[userID]
	if !exists {
		return nil, &ServiceError{Message: fmt.Sprintf("user with id %d not found", userID)}
	}

	actions, exists := s.db.Actions[userID]
	if !exists {
		return nil, &ServiceError{Message: fmt.Sprintf("no actions found for user id %d", userID)}
	}

	return &domain.TotalActions{Count: len(actions)}, nil
}
