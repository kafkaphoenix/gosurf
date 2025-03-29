package usecases

import (
	"fmt"
	"math"

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

// GetNextActionProbabilities return the next actions with their probabilities
// that could happen given an action type. Note: It the agregated users
// probability no per user.
func (s *ActionService) GetNextActionProbabilities(actionType string) (map[string]float64, error) {
	nextActionCounts := make(map[string]int)
	totalCount := 0

	// iterate over every user's action history (sorted by createdAt)
	for _, actions := range s.db.Actions {
		for i := range len(actions) - 1 { // iterate users actions
			if actions[i].Type == actionType { // match action type we are looking for
				nextAction := actions[i+1].Type // get next action type
				nextActionCounts[nextAction]++  // add or increase next action counter
				totalCount++
			}
		}
	}

	// best case, no user executed actions after our action type
	if totalCount == 0 {
		return map[string]float64{}, nil
	}

	// compute probabilities per next action given total actions done by user
	probabilities := make(map[string]float64)
	for nextAction, count := range nextActionCounts {
		probability := float64(count) / float64(totalCount)
		// we round to hundredths .XX as tenths add too much rounding error
		probabilities[nextAction] = math.Round(probability*100) / 100
	}

	return probabilities, nil
}
