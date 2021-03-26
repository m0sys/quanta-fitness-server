package userstore

import (
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type UserStore interface {
	Save(user entity.BaseUser) (entity.User, error)
	FindUserByUsername(username string) (entity.User, bool, error)
	FindUserByEmail(email string) (entity.User, bool, error)
}
