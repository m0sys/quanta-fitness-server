// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Auth struct {
	Token string `json:"token"`
}

type Login struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password string  `json:"password"`
}

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Confirm  string `json:"confirm"`
}

type NewWorkout struct {
	Title string `json:"title"`
}

type User struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Weight   float64 `json:"weight"`
	Height   float64 `json:"height"`
	Gender   string  `json:"gender"`
}

type Workout struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	User      *User     `json:"user"`
}

type WorkoutUpdate struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
