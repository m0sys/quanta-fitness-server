package entity

import (
	"time"
)

type BaseUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	BaseUser
	ID        int64     `json:"id"`
	Weight    float32   `json:"weight"`
	Height    float32   `json:"height"`
	Gender    string    `json:"gender"`
	Joined    time.Time `json:"joined"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRegister struct {
	BaseUser
	Confirm string `json:"confirm"`
}
