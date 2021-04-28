/*
Manager manages all personal use cases related to Athlete.
E.g. personal weight, height, etc...

@Author: mhd53
@Date: 2021/4/25
*/

package manager

import (
	"errors"
	"log"

	"github.com/mhd53/quanta-fitness-server/athlete"
)

// AthleteManager manages Athlete's personal data such as weight.
type AthleteManager interface {
	FetchWeightHistory() ([]WeightRecordRes, error)
	UpdateWeight(weight float64) error
	SetHeight(height float64) error
}

type manager struct {
	repo Repository
	ath  *athlete.Athlete
}

// NewManager create a new AthleteManager to manage `athlete`.
func NewManager(repository Repository, athlete *athlete.Athlete) AthleteManager {
	return &manager{
		repo: repository,
		ath:  athlete,
	}
}

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
