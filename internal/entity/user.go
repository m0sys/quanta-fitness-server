package entity

type BaseUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	BaseUser
	ID     int64   `json:"id"`
	Weight float32 `json:"weight"`
	Height float32 `json:"height"`
	Gender string  `json:"gender"`
}

type UserRegister struct {
	BaseUser
	Confirm string `json:"confirm"`
}
