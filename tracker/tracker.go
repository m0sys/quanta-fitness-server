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

func (t *tracker) AddExerciseToWorkoutLog(req AddExerciseToWorkoutLogReq) (ExerciseRes, error) {
	if t.wlog == nil {
		return ExerciseRes{}, errors.New("no WorkoutLog is assigned to Tracker")
	}
	if t.wlog.LogID != req.LogID {
		return ExerciseRes{}, errors.New("WorkoutLog does not match requested LogID")
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
		return ExerciseRes{}, errors.New("Internal Error")

	}

	return exerciseRes, nil
}
