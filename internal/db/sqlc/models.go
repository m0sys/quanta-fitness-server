// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"

	"github.com/google/uuid"
)

type Athlete struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Exercise struct {
	ID           uuid.UUID `json:"id"`
	Wpid         uuid.UUID `json:"wpid"`
	Aid          uuid.UUID `json:"aid"`
	Name         string    `json:"name"`
	TargetRep    int32     `json:"target_rep"`
	NumSets      int32     `json:"num_sets"`
	Weight       float64   `json:"weight"`
	RestDuration float64   `json:"rest_duration"`
	Pos          int32     `json:"pos"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ExerciseLog struct {
	ID           uuid.UUID `json:"id"`
	Wlid         uuid.UUID `json:"wlid"`
	Name         string    `json:"name"`
	TargetRep    int32     `json:"target_rep"`
	NumSets      int32     `json:"num_sets"`
	Weight       float64   `json:"weight"`
	RestDuration float64   `json:"rest_duration"`
	Completed    bool      `json:"completed"`
	Pos          int32     `json:"pos"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type SetLog struct {
	ID             uuid.UUID `json:"id"`
	Elid           uuid.UUID `json:"elid"`
	ActualRepCount int32     `json:"actual_rep_count"`
	Duration       float64   `json:"duration"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type WorkoutLog struct {
	ID         uuid.UUID `json:"id"`
	Aid        uuid.UUID `json:"aid"`
	Title      string    `json:"title"`
	Date       time.Time `json:"date"`
	CurrentPos int32     `json:"current_pos"`
	Completed  bool      `json:"completed"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type WorkoutPlan struct {
	ID        uuid.UUID `json:"id"`
	Aid       uuid.UUID `json:"aid"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
