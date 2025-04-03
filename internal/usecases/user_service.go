package usecases

import (
	"fmt"

	"github.com/kafkaphoenix/gosurf/internal/domain"
)

type UserService struct {
	db domain.DBRepo
}

func NewUserService(db domain.DBRepo) *UserService {
	return &UserService{db: db}
}

// GetUserByID return a user if exists.
func (s *UserService) GetUserByID(uid int) (*domain.User, error) {
	user, exists := s.db.GetUser(uid)
	if !exists {
		return nil, &ServiceError{Message: fmt.Sprintf("user with id %d not found", uid)}
	}

	return &user, nil
}

// GetTotalActionsByID return total actions of a user if exists.
func (s *UserService) GetTotalActionsByID(uid int) (*domain.TotalActions, error) {
	// we repeat code here because GetUserByID could change in the future
	// we could also guess the given user exists
	_, exists := s.db.GetUser(uid)
	if !exists {
		return nil, &ServiceError{Message: fmt.Sprintf("user with id %d not found", uid)}
	}

	actions, exists := s.db.GetActions(uid)
	if !exists {
		return nil, &ServiceError{Message: fmt.Sprintf("no actions found for user id %d", uid)}
	}

	return &domain.TotalActions{Count: len(actions)}, nil
}

// GetReferralIndex given all users and their referrals
// each node(user) and edge (referral) is visited once
// Complexity O(N*E). Space Complexity O(N).
func (s *UserService) GetReferralIndex() domain.ReferralIndex {
	users := s.db.GetAllUsers()
	// We are guessing worst case for capacity one user referred everyone
	referralIndex := make(domain.ReferralIndex, len(users))

	// BFS (Breadth-first search) to get each user's referrals
	for uid := range users {
		visited := make(map[int]bool)
		visited[uid] = true

		queue := []int{}
		// we avoid referrals being nil
		if referrals, exists := s.db.GetReferrals(uid); exists {
			queue = append(queue, referrals...)
		}

		count := 0

		for len(queue) > 0 {
			// pop from the front of the queue
			current := queue[0]
			queue = queue[1:]

			if !visited[current] {
				visited[current] = true
				count++
				// add their referrals to the queue if not visited before
				if referrals, exists := s.db.GetReferrals(current); exists {
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
