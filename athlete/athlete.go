// Athlete contains the Athlete entity.

package athlete

import (
	"errors"
	"math"
	"time"

	"github.com/mhd53/quanta-fitness-server/pkg/uuid"
	wl "github.com/mhd53/quanta-fitness-server/workoutlog"
)

// NOTE: All units will be in SI units.

// Athlete entity for representing what an athlete is.
type Athlete struct {
	uuid          string
	height        float64 // in meters
	weightHistory []WeightRecord
	workoutLogs   []wl.WorkoutLog
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

func (a *Athlete) WorkoutLogs() []wl.WorkoutLog {
	tmp := make([]wl.WorkoutLog, len(a.workoutLogs))
	copy(tmp, a.workoutLogs)
	return tmp
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

// AddWorkoutLog add WorkoutLog to Athlete.
func (a *Athlete) AddWorkoutLog(log wl.WorkoutLog) error {
	for _, l := range a.workoutLogs {
		if l.LogID() == log.LogID() {
			return errors.New("Workout Log is already assigned to Athlete")
		}
	}

	a.workoutLogs = append(a.workoutLogs, log)
	return nil
}

// RemoveWorkoutLog remove WorkoutLog from Athlete.
func (a *Athlete) RemoveWorkoutLog(log wl.WorkoutLog) error {
	deleted := false

	for i, l := range a.workoutLogs {
		if l.LogID() == log.LogID() {
			a.workoutLogs = removeWorkoutLog(a.workoutLogs, i)
			deleted = true
		}
	}

	if !deleted {
		return errors.New("WorkoutLog not found")
	}
	return nil
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

func removeWorkoutLog(slice []wl.WorkoutLog, idx int) []wl.WorkoutLog {
	return append(slice[:idx], slice[idx+1:]...)
}
