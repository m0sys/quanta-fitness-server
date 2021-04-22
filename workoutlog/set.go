package workoutlog

import (
	"errors"

	"github.com/mhd53/quanta-fitness-server/pkg/uuid"
)

// Set entity for representing what a set is.
type Set struct {
	SetID          string
	ActualRepCount int
}

// NewSet create a new Set.
func NewSet(actualRepCount int) (Set, error) {
	if err := validateActualRepCount(actualRepCount); err != nil {
		return &Set{}, err
	}

	return &Set{
		SetID:          uuid.GenerateUUID(),
		ActualRepCount: actualRepCount,
	}
}

// EditSet edit fields of Set.
func (s *Set) EditSet(actualRepCount int) error {
	if err := validateActualRepCount(actualRepCount); err != nil {
		return err
	}

	s.ActualRepCount = actualRepCount
	return nil
}

func validateActualRepCount(actualRepCount int) error {
	if actualRepCount < 0 {
		return errors.New("actual rep count must be a positive number")
	}
	return nil
}
