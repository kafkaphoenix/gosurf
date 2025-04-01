package db

import (
	"encoding/json"
	"os"
	"slices"

	"github.com/kafkaphoenix/gosurf/internal/domain"
)

type FakeDB struct {
	Users         map[int]domain.User
	Actions       map[int][]domain.Action
	ReferralGraph map[int][]int
}

// NewFakeDB initializes the fake db by loading JSON data.
func NewFakeDB(userFile, actionFile string) (*FakeDB, error) {
	userData, err := os.ReadFile(userFile)
	if err != nil {
		return nil, &DBError{Message: "error reading user file", Err: err}
	}

	var userList []domain.User
	if err := json.Unmarshal(userData, &userList); err != nil {
		return nil, &DBError{Message: "error unmarshaling user data", Err: err}
	}

	users := make(map[int]domain.User, len(userList))
	for _, user := range userList {
		users[user.ID] = user
	}

	actionData, err := os.ReadFile(actionFile)
	if err != nil {
		return nil, &DBError{Message: "error reading action file", Err: err}
	}

	var actionList []domain.Action
	if err := json.Unmarshal(actionData, &actionList); err != nil {
		return nil, &DBError{Message: "error unmarshaling action data", Err: err}
	}

	actions := make(map[int][]domain.Action, len(actionList))
	// worst case every action is a REFER_USER action
	referralGraph := make(map[int][]int, len(actionList))

	// we group actions by user id
	for _, action := range actionList {
		actions[action.UserID] = append(actions[action.UserID], action)

		// build referral graph
		if action.Type == "REFER_USER" {
			referralGraph[action.UserID] = append(referralGraph[action.UserID], action.TargetUser)
		}
	}

	// sort actions for each user by CreatedAt
	// complexity O(N Log N)
	for uid := range actions {
		slices.SortFunc(actions[uid], func(a, b domain.Action) int {
			if a.CreatedAt.Before(b.CreatedAt) {
				return -1
			}

			if a.CreatedAt.After(b.CreatedAt) {
				return 1
			}

			return 0
		})
	}

	return &FakeDB{Users: users, Actions: actions, ReferralGraph: referralGraph}, nil
}
