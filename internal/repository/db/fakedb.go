package db

import (
	"encoding/json"
	"os"
	"slices"
	"sync"

	"github.com/kafkaphoenix/gosurf/internal/domain"
)

// map: fast lookups by user id <- our case
// slice: faster iteration
type FakeDB struct {
	Users map[int]domain.User
	// not created domain ActionList as we just return the count
	// in this demo
	Actions       map[int][]domain.Action
	ReferralGraph map[int][]int
	// mu            sync.RWMutex // mutex for concurrent access
}

// NewFakeDB initialize the fake db by loading concurrently JSON data.
func NewFakeDB(userFile, actionFile string) (*FakeDB, error) {
	db := &FakeDB{}

	var wg sync.WaitGroup

	wg.Add(2)

	var userErr, actionErr error

	go func() {
		defer wg.Done()

		userErr = db.loadUsers(userFile)
	}()

	go func() {
		defer wg.Done()

		actionErr = db.loadActions(actionFile)
	}()

	wg.Wait()

	if userErr != nil {
		return nil, &FakeDBError{Message: "error loading users", Err: userErr}
	}

	if actionErr != nil {
		return nil, &FakeDBError{Message: "error loading actions", Err: actionErr}
	}

	return db, nil
}

// loadUsers load user data from a JSON file.
func (db *FakeDB) loadUsers(userFile string) error {
	userData, err := os.ReadFile(userFile)
	if err != nil {
		// it could be argued here that 'error' should not be added in an error message,
		// only in the root of the error chain to avoid duplicating error 50-50
		return &FakeDBError{Message: "error reading user file", Err: err}
	}

	var users []domain.User
	if err := json.Unmarshal(userData, &users); err != nil {
		return &FakeDBError{Message: "error unmarshaling user data", Err: err}
	}

	// not required as loading(creation) happens one time at the start
	// but essential in a real db for consistency
	// db.mu.Lock()
	// defer db.mu.Unlock()

	// pre-allocate for efficiency
	db.Users = make(map[int]domain.User, len(users))
	for _, u := range users {
		// careful with partial updates in a real db
		db.Users[u.ID] = u
	}

	return nil
}

// loadActions load action data from a JSON file and build the referral graph.
func (db *FakeDB) loadActions(actionFile string) error {
	actionData, err := os.ReadFile(actionFile)
	if err != nil {
		return &FakeDBError{Message: "error reading action file", Err: err}
	}

	var actions []domain.Action
	if err := json.Unmarshal(actionData, &actions); err != nil {
		return &FakeDBError{Message: "error unmarshaling action data", Err: err}
	}

	db.Actions = make(map[int][]domain.Action, len(actions))
	// worst case every action is a REFER_USER action
	db.ReferralGraph = make(map[int][]int, len(actions))

	// group actions by user id
	for _, a := range actions {
		db.Actions[a.UserID] = append(db.Actions[a.UserID], a)

		// build referral graph
		if a.Type == "REFER_USER" {
			db.ReferralGraph[a.UserID] = append(db.ReferralGraph[a.UserID], a.TargetUser)
		}
	}

	// sort actions for each user by CreatedAt
	// complexity O(N Log N)
	for uid := range db.Actions {
		// modify the underlying slice
		slices.SortFunc(db.Actions[uid], func(a, b domain.Action) int {
			if a.CreatedAt.Before(b.CreatedAt) {
				return -1
			}

			if a.CreatedAt.After(b.CreatedAt) {
				return 1
			}

			return 0
		})
	}

	return nil
}

// GetUser return a user by ID.
func (db *FakeDB) GetUser(id int) (domain.User, bool) {
	// not required as loading(creation) happens one time at the start
	// and not updates in the demo
	// but essential in a real db for consistency
	// db.mu.RLock()
	// defer db.mu.RUnlock()
	user, exists := db.Users[id]
	return user, exists
}

// GetAllUsers return all users.
func (db *FakeDB) GetAllUsers() map[int]domain.User {
	return db.Users
}

// GetActions return all actions for a specific user.
func (db *FakeDB) GetActions(uid int) ([]domain.Action, bool) {
	actions, exists := db.Actions[uid]
	return actions, exists
}

// GetAllActions return all actions grouped by user.
func (db *FakeDB) GetAllActions() map[int][]domain.Action {
	return db.Actions
}

// GetReferrals return the list of users referred by a specific user.
func (db *FakeDB) GetReferrals(uid int) ([]int, bool) {
	referees, exists := db.ReferralGraph[uid]
	return referees, exists
}
