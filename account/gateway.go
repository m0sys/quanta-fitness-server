package account

type AccountGateway interface {
	SignUp(uname, password, email string) error
	FindUserByUsername(uname string) (UserResponse, bool)
}
