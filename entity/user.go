package entity

type User struct {
	ID       int64   `json:"id"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Weight   float32 `json:"weight"`
	Height   float32 `json:"height"`
	Gender   string  `json:"gender"`
}

type UserPublic struct {
	Username string  `json:"username"`
	Weight   float32 `json:"weight"`
	Height   float32 `json:"height"`
	Gender   string  `json:"gender"`
}

type UsernameLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type EmailLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegister struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Confirm  string `json:"confirm"`
}
