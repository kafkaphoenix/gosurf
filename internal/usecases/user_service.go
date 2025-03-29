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

// Complexity O(N**2 + N*E). Space Complexity O(N).
func (s *UserService) GetReferralIndex() domain.ReferralIndex {
	// We are guessing worst case for capacity one user referred everyone
	referralIndex := make(domain.ReferralIndex, len(s.db.Users)-1)

	// BFS (Breadth-first search) to get each user's referrals
	for uid := range s.db.Users {
		visited := make(map[int]bool)
		visited[uid] = true
		queue := s.db.ReferralGraph[uid]
		count := 0

		for len(queue) > 0 {
			// pop from the front of the queue
			current := queue[0]
			queue = queue[1:]

			if !visited[current] {
				visited[current] = true
				count++
				// add their referrals to the queue if not visited before
				if referrals, ok := s.db.ReferralGraph[current]; ok {
					for _, referral := range referrals {
						if !visited[referral] {
							queue = append(queue, referral)
						}
					}
				}
			}
		}

		referralIndex[uid] = domain.Referral{Count: count}
	}

	return referralIndex
}
