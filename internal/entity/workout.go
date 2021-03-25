package entity

import (
	"time"
)

type BaseWorkout struct {
	Title    string `json:"title"`
	Username string `json:"username"`
}

type Workout struct {
	BaseWorkout
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
