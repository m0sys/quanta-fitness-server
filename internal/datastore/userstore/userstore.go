package userstore

import (
	"github.com/m0sys/quanta-fitness-server/internal/entity"
)

type UserStore interface {
	Save(user entity.BaseUser) (entity.User, error)
	FindUserByUsername(username string) (entity.User, bool, error)
	FindUserByEmail(email string) (entity.User, bool, error)
	DeleteUser(id int64) (bool, error)
}
