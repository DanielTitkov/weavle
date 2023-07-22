package domain

import (
	"time"

	"github.com/google/uuid"
)

type Story struct {
	ID        uuid.UUID
	Status    string
	Sentences []*Sentence
}

type Sentence struct {
	ID      uuid.UUID // sentence in domain logic always goes inside a story. It doesn't need an id here.
	Text    string
	Author  *User
	StoryID uuid.UUID
}

type User struct {
	ID        uuid.UUID
	IPAddress string
}

// system types
type (
	SystemSummary struct {
		ID            int
		Users         int
		Tests         int
		FinishedTakes int
		Responses     int
		CreateTime    time.Time
	}
	Event struct {
		Name      string
		StartTime time.Time
		EndTime   time.Time
		Elapsed   time.Duration
	}
)
