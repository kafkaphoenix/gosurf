package usecases

import (
	"fmt"

	"github.com/kafkaphoenix/gosurf/internal/domain"
	"github.com/kafkaphoenix/gosurf/internal/repository"
)

type ActionService struct {
	db *repository.FakeDB
}

func NewActionService(db *repository.FakeDB) *ActionService {
	return &ActionService{db: db}
}

// GetTotalActionsByID return total actions of a user if exists.
func (s *ActionService) GetTotalActionsByID(userID int) (*domain.TotalActions, error) {
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
