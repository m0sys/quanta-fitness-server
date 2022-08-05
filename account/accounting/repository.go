package accounting

import "github.com/m0sys/quanta-fitness-server/account/user"

type Repository interface {
	FindUserByUname(uname string) (user.User, bool, error)
	FindUserByEmail(email string) (user.User, bool, error)
	StoreUser(newUser user.User) error
}
