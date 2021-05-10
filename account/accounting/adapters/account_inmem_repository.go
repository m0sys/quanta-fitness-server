package adapters

import (
	"time"

	a "github.com/mhd53/quanta-fitness-server/account/accounting"
	"github.com/mhd53/quanta-fitness-server/account/user"
)

type repo struct {
	users map[string]inRepoUser
}

func NewInMemRepo() a.Repository {
	return &repo{
		users: make(map[string]inRepoUser),
	}
}

func (r *repo) FindUserByUname(uname string) (user.User, bool, error) {
	data, ok := r.users[uname]
	if !ok {
		return user.User{}, false, nil
	}

	u, err := user.RestoreUser(data.Username, data.Email, data.HashedPwd, data.Joined)
	if err != nil {
		return user.User{}, false, err
	}

	return u, true, nil
}

func (r *repo) FindUserByEmail(email string) (user.User, bool, error) {
	for _, val := range r.users {
		if val.Email == email {
			u, err := user.RestoreUser(val.Username, val.Email, val.HashedPwd, val.Joined)
			if err != nil {
				return user.User{}, false, err
			}

			return u, true, nil
		}
	}
	return user.User{}, false, nil
}

func (r *repo) StoreUser(newUser user.User) error {
	data := inRepoUser{
		Username:  newUser.Username(),
		Email:     newUser.Email(),
		HashedPwd: newUser.HashedPwd(),
		Joined:    newUser.Joined(),
		UpdatedAt: newUser.Joined(),
	}

	r.users[newUser.Username()] = data
	return nil

}

type inRepoUser struct {
	Username  string
	Email     string
	HashedPwd string
	Joined    time.Time
	UpdatedAt time.Time
}
