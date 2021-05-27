package managing

import (
	"github.com/mhd53/quanta-fitness-server/account/user"
	"github.com/mhd53/quanta-fitness-server/manager/athlete"
)

// Repository repo for persisting all Athlete related data.
type Repository interface {
	StoreAthlete(usr user.User, ath athlete.Athlete) error
	FindAthleteByUname(usr user.User) (athlete.Athlete, bool, error)
	FindAthleteByID(id string) (athlete.Athlete, bool, error)
}
