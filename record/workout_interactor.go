package record

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/mhd53/quanta-fitness-server/internal/id"
	w "github.com/mhd53/quanta-fitness-server/workout"
)

type workoutInteractor interface {
	CreateWorkout(title string, date time.Time, uname string) (w.Workout, error)
	UpdateWorkout(id id.ID, uname, title string, date time.Time) error
}

func (*interactor) CreateWorkout(title string, date time.Time, uname string) (w.Workout, error) {
	user, ok := aGateway.FindUserByUsername(uname)
	if !ok {
		return w.Workout{}, fmt.Errorf("couldn't create workout: %w", errAccessDenied)
	}

	workout, err := w.NewWorkout(title, date)
	if err != nil {
		return w.Workout{}, err
	}

	athlete, ok := tGateway.FindAtheleteByUname(uname)
	if !ok {
		log.Fatal("Error: User is not an athlete!?")
	}

	workout.AID = athlete.AID

	err := tGateway.SaveWorkout(workout)
	if err != nil {
		return w.Workout{}, errors.New("Couldn't save user!")
	}

	return workout, nil
}

func (*interactor) UpdateWorkout(id id.ID, uname, title string, date time.Time) error {
	user, ok := aGateway.FindUserByUsername(uname)
	if !ok {
		return w.Workout{}, fmt.Errorf("couldn't update workout: %w", errAccessDenied)
	}

	prev, ok := wGateway.FindWorkoutByID(id)
	if !ok {
		return fmt.Errorf("couldn't update workout: %w", errWorkoutNotFound)
	}

	updated := w.Workout{
		ID:    prev.ID,
		AID:   prev.AID,
		Title: title,
		Date:  date,
	}

}
