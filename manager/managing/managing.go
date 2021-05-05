/*
Manager manages all personal use cases related to Athlete.
E.g. personal weight, height, etc...

@Author: mhd53
@Date: 2021/4/25
*/

package managing

import (
	"log"

	"github.com/mhd53/quanta-fitness-server/account/user"
	"github.com/mhd53/quanta-fitness-server/manager/athlete"
)

// AthleteManager manages Athlete's personal data such as weight.
type ManagingService struct {
	repo Repository
}

// NewManager create a new AthleteManager to manage `athlete`.
func NewManagingService(repository Repository) ManagingService {
	return ManagingService{
		repo: repository,
	}
}

func (m ManagingService) CreateNewAthlete(usr user.User) (athlete.Athlete, error) {
	ath := athlete.NewAthlete()
	err := m.repo.StoreAthlete(usr, ath)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return athlete.Athlete{}, errInternal
	}

	return ath, nil
}

func (m ManagingService) FetchAthlete(usr user.User) (athlete.Athlete, error) {
	ath, found, err := m.repo.FindAthleteByUname(usr)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return athlete.Athlete{}, errInternal
	}

	if !found {
		return athlete.Athlete{}, ErrAthleteNotFound
	}

	return ath, nil
}
