package adapters

import (
	"time"

	"github.com/mhd53/quanta-fitness-server/account/user"
	"github.com/mhd53/quanta-fitness-server/manager/athlete"
	m "github.com/mhd53/quanta-fitness-server/manager/managing"
)

type repo struct {
	athletes map[string]inRepoAthlete
}

func NewInMemRepo() m.Repository {
	return &repo{
		athletes: make(map[string]inRepoAthlete),
	}
}

func (r *repo) StoreAthlete(usr user.User, ath athlete.Athlete) error {
	now := time.Now()
	data := inRepoAthlete{
		ID:        ath.AthleteID(),
		Username:  usr.Username(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	r.athletes[ath.AthleteID()] = data
	return nil
}

func (r *repo) FindAthleteByUname(usr user.User) (athlete.Athlete, bool, error) {
	uname := usr.Username()
	for _, val := range r.athletes {
		if val.Username == uname {
			return athlete.RestoreAthlete(val.ID), true, nil
		}
	}

	return athlete.Athlete{}, false, nil
}

func (r *repo) FindAthleteByID(id string) (athlete.Athlete, bool, error) {
	for _, val := range r.athletes {
		if val.ID == id {
			return athlete.RestoreAthlete(val.ID), true, nil
		}
	}

	return athlete.Athlete{}, false, nil
}

type inRepoAthlete struct {
	ID        string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
