// Set contains the Set entity.

package workoutlog

import (
	"errors"

	"github.com/mhd53/quanta-fitness-server/pkg/uuid"
)

// Set entity for representing what a set is.
type Set struct {
	uuid           string
	actualRepCount int
}

// NewSet create a new Set.
func NewSet(actualRepCount int) (Set, error) {
	if err := validateActualRepCount(actualRepCount); err != nil {
		return Set{}, err
	}

	return Set{
		uuid:           uuid.GenerateUUID(),
		actualRepCount: actualRepCount,
	}, nil
}

func (s *Set) SetID() string {
	return s.uuid
}

func (s *Set) ActualRepCount() int {
	return s.actualRepCount
}

// EditSet edit fields of Set.
func (s *Set) EditSet(actualRepCount int) error {
	if err := validateActualRepCount(actualRepCount); err != nil {
		return err
	}

	s.actualRepCount = actualRepCount
	return nil
}

func validateActualRepCount(actualRepCount int) error {
	if actualRepCount < 0 {
		return errors.New("actual rep count must be a positive number")
	}
	return nil
}
