package domain

import "time"

// User represent basic user's information.
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// Action represent an action performed by a user.
type Action struct {
	ID         int       `json:"id"`
	Type       string    `json:"type"`
	UserID     int       `json:"userId"`
	TargetUser int       `json:"targetUser,omitempty"`
	CreatedAt  time.Time `json:"createdAt"`
}

type Actions []Action

type TotalActions struct {
	Count int `json:"count"`
}
