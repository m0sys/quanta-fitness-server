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

type tracker struct {
	repo Repository
	ath  *athlete.Athlete
}

// FIXME: Get rid of that pointer for ath.

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

	return repoWlog, nil
}

func (t *tracker) AddExerciseToWorkoutLog(req AddExerciseToWorkoutLogReq) (ExerciseRes, error) {
	wlogRes, found, err := t.repo.FindWorkoutLogByID(req.LogID)
	if err != nil {
		log.Printf("error while searching for WorkoutLog in repo: %s", err.Error())
		return ExerciseRes{}, errors.New("Internal Error")
	}

	if !found {
		return ExerciseRes{}, errors.New("Workout Log not found")
	}

	wlog := wl.WorkoutLog{
		LogID: wlogRes.LogID,
		Title: wlogRes.Title,
		Date:  wlogRes.Date,
	}

	exercisesRes, err := t.repo.FindAllExercisesForWorkoutLog(wlog)
	if err != nil {
		log.Printf("error while fetching all Exercises for WorkoutLog in repo: %s", err.Error())
		return ExerciseRes{}, errors.New("Internal Error")
	}

	var exercises []wl.Exercise
	for _, e := range exercisesRes {
		exercise := wl.Exercise{
			ExerciseID: e.ExerciseID,
			Name:       e.Name,
			Weight:     e.Weight,
			TargetRep:  e.TargetRep,
			RestTime:   e.RestTime,
		}
		exercises = append(exercises, exercise)
	}
	wlog.Exercises = exercises

	newExercise, err := wl.NewExercise(req.Name, req.Weight, req.RestTime, req.TargetRep, len(exercises))
	if err != nil {
		return ExerciseRes{}, err
	}

	err = wlog.AddExercise(newExercise)
	if err != nil {
		return ExerciseRes{}, err
	}

	exerciseRes, err := t.repo.AddExerciseToWorkoutLog(wlog, newExercise)
	if err != nil {
		log.Printf("error while adding Exercise to WorkoutLog in repo: %s", err.Error())
		return ExerciseRes{}, errors.New("Internal Error")

	}

	return exerciseRes, nil
}
