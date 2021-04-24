// Tracker contains all use cases for tracking Workouts for Athlete.

package tracker

import (
	"errors"
	"log"

	"github.com/mhd53/quanta-fitness-server/athlete"
	wl "github.com/mhd53/quanta-fitness-server/workoutlog"
)

// WorkoutTracker tracks Athlete's Workout and stores the data.
type WorkoutTracker interface {
	CreateWorkoutLog(title string) (WorkoutLogRes, error)
	AddExerciseToWorkoutLog(req AddExerciseToWorkoutLogReq) (ExerciseRes, error)
	AddSetToExercise(req AddSetToExerciseReq) (SetRes, error)
	SetWorkoutLog(id string) error
	RemoveExerciseFromWorkoutLog(req RemoveExerciseFromWorkoutLogReq) error
}

// FIXME: Get rid of that pointer for ath.
type tracker struct {
	repo Repository
	ath  *athlete.Athlete
	wlog *wl.WorkoutLog
}

// NewTracker create a new WorkoutTracker for Athlete.
func NewTracker(repository Repository, athlete *athlete.Athlete) WorkoutTracker {
	return &tracker{
		repo: repository,
		ath:  athlete,
	}
}

func (t *tracker) CreateWorkoutLog(title string) (WorkoutLogRes, error) {
	wlog, err := wl.NewWorkoutLog(title)
	if err != nil {
		return WorkoutLogRes{}, err
	}

	err = t.ath.AddWorkoutLog(wlog)
	if err != nil {
		log.Fatal("couldn't add newly created WorkoutLog to Athlete")
	}

	repoWlog, err := t.repo.StoreWorkoutLog(wlog, *t.ath)
	if err != nil {
		log.Printf("%s: couldn't store WorkoutLog: %s", "tracker", err.Error())
	}

	t.wlog = &wlog
	return repoWlog, nil
}

func (t *tracker) SetWorkoutLog(id string) error {
	return nil
}

func (t *tracker) AddExerciseToWorkoutLog(req AddExerciseToWorkoutLogReq) (ExerciseRes, error) {
	if t.wlog == nil {
		return ExerciseRes{}, errNilWorkoutLog
	}
	if t.wlog.LogID != req.LogID {
		return ExerciseRes{}, errLogIDMismatch
	}

	newExercise, err := wl.NewExercise(req.Name, req.Weight, req.RestTime, req.TargetRep, len(t.wlog.Exercises))
	if err != nil {
		return ExerciseRes{}, err
	}

	err = t.wlog.AddExercise(newExercise)
	if err != nil {
		return ExerciseRes{}, err
	}

	exerciseRes, err := t.repo.AddExerciseToWorkoutLog(*t.wlog, newExercise)
	if err != nil {
		log.Printf("error while adding Exercise to WorkoutLog in repo: %s", err.Error())
		return ExerciseRes{}, errInternal

	}

	return exerciseRes, nil
}

func (t *tracker) AddSetToExercise(req AddSetToExerciseReq) (SetRes, error) {
	if t.wlog == nil {
		return SetRes{}, errNilWorkoutLog
	}
	if t.wlog.LogID != req.LogID {
		return SetRes{}, errLogIDMismatch
	}

	found := false
	var exercise wl.Exercise
	for _, e := range t.wlog.Exercises {
		if e.ExerciseID == req.ExerciseID {
			found = true
			exercise = e
		}
	}

	if !found {
		return SetRes{}, errors.New("Exercise not found")
	}

	newSet, err := wl.NewSet(req.ActualRepCount)
	if err != nil {
		return SetRes{}, err
	}

	err = exercise.AddSet(newSet)
	if err != nil {
		log.Fatal("couldn't add newly created Set to Exercise")
	}

	setRes, err := t.repo.AddSetToExercise(exercise, newSet)
	if err != nil {
		log.Printf("error while adding Set to Exercise in repo: %s", err.Error())
		return SetRes{}, errInternal
	}

	return setRes, nil
}

func (t *tracker) RemoveExerciseFromWorkoutLog(req RemoveExerciseFromWorkoutLogReq) error {
	return nil
}
