package userstore

import (
	"errors"
	"time"

	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type store struct {
	users  map[string]entity.User
	lastID int64
}

func NewMockUserStore() UserStore {
	return &store{users: make(map[string]entity.User)}
}

func (s *store) Save(newUser entity.BaseUser) (entity.User, error) {
	created := entity.User{
		BaseUser:  newUser,
		ID:        s.lastID,
		Weight:    0,
		Height:    0,
		Gender:    "N/A",
		Joined:    time.Now(),
		UpdatedAt: time.Now(),
	}
	s.users[newUser.Username] = created
	s.lastID += 1

	return created, nil
}

func (s *store) FindUserByUsername(uname string) (entity.User, bool, error) {
	found := s.users[uname]

	if isEmpty(found) {
		return entity.User{}, false, nil
	}

	return found, true, nil
}

func (s *store) FindUserByEmail(email string) (entity.User, bool, error) {
	var found entity.User

	for k := range s.users {
		entry := s.users[k]
		if entry.Email == email {
			found = entry
		}
	}

	if isEmpty(found) {
		return entity.User{}, false, nil
	}

	return found, true, nil
}

func (s *store) DeleteUser(id int64) (bool, error) {
	return false, errors.New("Not Implemented!")
}

func isEmpty(found entity.User) bool {
	return found == (entity.User{})
}
