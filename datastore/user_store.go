package datastore

import (
	"github.com/mhd53/quanta-fitness-server/entity"
)

type UserStore interface {
	Save(user *entity.UserRegister) (*entity.User, error)
	FindUserByUsername(username string) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
}
