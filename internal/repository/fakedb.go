package repository

import (
	"encoding/json"
	"os"

	"github.com/kafkaphoenix/gosurf/internal/domain"
)

type FakeDB struct {
	Users   map[int]domain.User
	Actions map[int][]domain.Action
}

// NewFakeDB initializes the fake db by loading JSON data.
func NewFakeDB(userFile, actionFile string) (*FakeDB, error) {
	users := make(map[int]domain.User)
	actions := make(map[int][]domain.Action)

	userData, err := os.ReadFile(userFile)
	if err != nil {
		return nil, err
	}

	var userList []domain.User
	if err := json.Unmarshal(userData, &userList); err != nil {
		return nil, err
	}

	for _, user := range userList {
		users[user.ID] = user
	}

	// Load actions
	actionData, err := os.ReadFile(actionFile)
	if err != nil {
		return nil, err
	}

	var actionList []domain.Action
	if err := json.Unmarshal(actionData, &actionList); err != nil {
		return nil, err
	}

	for _, action := range actionList {
		actions[action.UserID] = append(actions[action.UserID], action)
	}

	return &FakeDB{Users: users, Actions: actions}, nil
}
