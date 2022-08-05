// Athlete contains the Athlete entity.

package athlete

import (
	"errors"
	"math"
	"time"

	"github.com/m0sys/quanta-fitness-server/pkg/uuid"
)

// NOTE: All units will be in SI units.

// Athlete entity for representing what an athlete is.
type Athlete struct {
	uuid          string
	height        float64 // in meters
	weightHistory []WeightRecord
}

// WeightRecord holds recording for Athlete's weight.
type WeightRecord struct {
	amount float64 // in kg
	date   time.Time
}

func (wr *WeightRecord) Amount() float64 {
	return wr.amount
}

func (wr *WeightRecord) Date() time.Time {
	return wr.date
}

// NewAthlete create a new Athlete.
func NewAthlete() Athlete {
	return Athlete{
		uuid: uuid.GenerateUUID(),
	}
}

// This should only be used to restore Athlete from persistence layer.
func RestoreAthlete(id string) Athlete {
	return Athlete{
		uuid: id,
	}
}

func (a *Athlete) AthleteID() string {
	return a.uuid
}

func (a *Athlete) Height() float64 {
	return a.height
}

// SetHeight set the height of Athlete.
func (a *Athlete) SetHeight(height float64) error {
	if err := validateHeight(height); err != nil {
		return err
	}
	a.height = height
	return nil
}

// UpdateWeight update the weight of Athlete.
func (a *Athlete) UpdateWeight(weight float64) (WeightRecord, error) {
	if err := validateWeight(weight); err != nil {
		return WeightRecord{}, err
	}

	newWeight := WeightRecord{
		amount: roundToTwoDecimalPlaces(weight),
		date:   time.Now(),
	}

	a.weightHistory = append(a.weightHistory, newWeight)
	return newWeight, nil
}

// Helper funcs.

func validateHeight(height float64) error {
	if height < 0 {
		return errors.New("height must be a positive number")
	}
	return nil
}

func validateWeight(weight float64) error {
	if weight < 0 {
		return errors.New("weight must be a positive number")
	}
	return nil
}

func roundToTwoDecimalPlaces(num float64) float64 {
	return math.Round(num*100) / 100
}
