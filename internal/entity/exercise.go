package entity

import (
	"time"
)

type BaseExercise struct {
	Name     string `json:"name"`
	WID      string `json:"wid"`
	Username string `json:"username"`
}

// NOTE: All units will be in SI units.
type Metrics struct {
	Weight    float32 `json:"weight"`
	TargetRep int     `json:"target_rep"`
	RestTime  float32 `json:"rest_time"`
	NumSets   int     `json:"num_sets"`
}

type Exercise struct {
	BaseExercise
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Metrics
}

type ExerciseUpdate struct {
	Name string `json:"name"`
	Metrics
}
