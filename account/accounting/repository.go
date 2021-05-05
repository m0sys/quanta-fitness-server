package accounting

import "github.com/mhd53/quanta-fitness-server/account/user"

type Repository interface {
	FindUserByUname(uname string) (user.User, bool, error)
	FindUserByEmail(email string) (user.User, bool, error)
	StoreUser(newUser user.User) error
}
