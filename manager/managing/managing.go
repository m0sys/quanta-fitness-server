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

/*
func (m *manager) FetchWeightHistory() ([]WeightRecordRes, error) {
	var history []WeightRecordRes
	history, err := m.repo.FindAllWeightRecords(m.ath.AthleteID())
	if err != nil {
		log.Printf("%s: couldn't fetch all Weight Records from repo: %s", "manager", err.Error())
		return history, errors.New("Internal error")

	}
	return history, nil
}

func (m *manager) UpdateWeight(weight float64) error {
	record, err := m.ath.UpdateWeight(weight)
	if err != nil {
		return err
	}

	return m.repo.StoreWeightRecord(m.ath.AthleteID(), record)
}

func (m *manager) SetHeight(height float64) error {
	return nil
}
*/
