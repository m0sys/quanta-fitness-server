package entity

import (
	"time"
)

type BaseExercise struct {
	Name     string `json:"name"`
	WID      int64  `json:"wid"`
	Username int64  `json:"username"`
}

type Exercise struct {
	BaseExercise
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Weight    float32   `json:"weight"`
	TargetRep int       `json:"target_rep"`
	RestTime  float32   `json:"rest_time"`
	NumSets   int       `json:"num_sets"`
}