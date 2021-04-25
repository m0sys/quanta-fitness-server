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
	RemoveExerciseFromWorkoutLog(exerciseID string) error
	RemoveSetFromExercise(setID string, exerciseID string) error
	EditWorkoutLog(req EditWorkoutLogReq) (WorkoutLogRes, error)
	EditExercise(req EditExerciseReq) (ExerciseRes, error)
	EditSet(req EditSetReq) (SetRes, error)
	FetchAllWorkoutLogs() ([]WorkoutLogRes, error)
	FetchAllExercisesForWorkoutLog() ([]ExerciseRes, error)
	FetchAllSetsForExercise(exerciseID string) ([]SetRes, error)
}

// FIXME: Get rid of that pointer for ath.
// Precondition: WorkoutTracker will only retrieve data pertaining `wlog` and `ath`.
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

	// NOTE: maybe we don't need to check for LogID since we want to add to
	//		 the WorkoutLog that is currently in wlog.
	if t.wlog.LogID != req.LogID {
		return SetRes{}, errLogIDMismatch
	}

	found := false
	var foundIdx int
	var exercise wl.Exercise
	for i, e := range t.wlog.Exercises {
		if e.ExerciseID == req.ExerciseID {
			found = true
			foundIdx = i
			exercise = e
		}
	}

	if !found {
		return SetRes{}, errExerciseNotFound
	}

	newSet, err := wl.NewSet(req.ActualRepCount)
	if err != nil {
		return SetRes{}, err
	}
	err = exercise.AddSet(newSet)
	if err != nil {
		log.Fatal("couldn't add newly created Set to Exercise")
	}

	t.wlog.Exercises[foundIdx] = exercise

	setRes, err := t.repo.AddSetToExercise(exercise, newSet)
	if err != nil {
		log.Printf("error while adding Set to Exercise in repo: %s", err.Error())
		return SetRes{}, errInternal
	}

	return setRes, nil
}

func (t *tracker) RemoveExerciseFromWorkoutLog(exerciseID string) error {
	if t.wlog == nil {
		return errNilWorkoutLog
	}

	found := false
	var exercise wl.Exercise
	for _, e := range t.wlog.Exercises {
		if e.ExerciseID == exerciseID {
			found = true
			exercise = e
		}
	}

	if !found {
		return errExerciseNotFound
	}

	err := t.wlog.RemoveExercise(exercise)
	if err != nil {
		return err
	}

	err = t.repo.DeleteExercise(exerciseID)
	if err != nil {
		log.Printf("%s: couldn't delete Exercise from repo: %s", "tracker", err.Error())
		return errInternal
	}

	return nil
}
func (t *tracker) RemoveSetFromExercise(setID string, exerciseID string) error {
	if t.wlog == nil {
		return errNilWorkoutLog
	}

	found := false
	var exercise wl.Exercise
	for _, e := range t.wlog.Exercises {
		if e.ExerciseID == exerciseID {
			found = true
			exercise = e
		}
	}

	if !found {
		return errExerciseNotFound
	}

	found = false
	var set wl.Set
	for _, s := range exercise.Sets {
		if s.SetID == setID {
			found = true
			set = s
		}
	}

	if !found {
		return errors.New("Set not found")
	}

	err := exercise.RemoveSet(set)
	if err != nil {
		return err
	}

	err = t.repo.DeleteSet(set.SetID)

	if err != nil {
		log.Printf("%s: couldn't delete Set from repo: %s", "tracker", err.Error())
		return errInternal
	}

	return nil
}

func (t *tracker) EditWorkoutLog(req EditWorkoutLogReq) (WorkoutLogRes, error) {
	if t.wlog == nil {
		return WorkoutLogRes{}, errNilWorkoutLog
	}
	if t.wlog.LogID != req.LogID {
		return WorkoutLogRes{}, errLogIDMismatch
	}

	err := t.wlog.EditWorkoutLog(req.Title, req.Date)
	if err != nil {
		return WorkoutLogRes{}, err
	}

	res, err := t.repo.UpdateWorkoutLog(req)
	if err != nil {
		log.Printf("%s: couldn't update WorkoutLog from repo: %s", "tracker", err.Error())
		return WorkoutLogRes{}, errInternal
	}

	return res, nil
}

func (t *tracker) EditExercise(req EditExerciseReq) (ExerciseRes, error) {
	if t.wlog == nil {
		return ExerciseRes{}, errNilWorkoutLog
	}

	found := false
	var foundIdx int
	var exercise wl.Exercise
	for i, e := range t.wlog.Exercises {
		if e.ExerciseID == req.ExerciseID {
			found = true
			foundIdx = i
			exercise = e
		}
	}

	if !found {
		return ExerciseRes{}, errExerciseNotFound
	}

	err := exercise.EditExercise(req.Name, req.Weight, req.RestTime, req.TargetRep)
	if err != nil {
		return ExerciseRes{}, err
	}

	t.wlog.Exercises[foundIdx] = exercise

	res, err := t.repo.UpdateExercise(req)
	if err != nil {
		log.Printf("%s: couldn't update Exercise from repo: %s", "tracker", err.Error())
		return ExerciseRes{}, errInternal

	}

	return res, nil
}

func (t *tracker) EditSet(req EditSetReq) (SetRes, error) {
	if t.wlog == nil {
		return SetRes{}, errNilWorkoutLog
	}

	found := false
	var foundIdxE int
	var exercise wl.Exercise
	for i, e := range t.wlog.Exercises {
		if e.ExerciseID == req.ExerciseID {
			found = true
			foundIdxE = i
			exercise = e
		}
	}

	if !found {
		return SetRes{}, errExerciseNotFound
	}
	log.Print(foundIdxE)

	found = false
	var foundIdxS int
	var set wl.Set
	for i, s := range exercise.Sets {
		if s.SetID == req.SetID {
			found = true
			foundIdxS = i
			set = s
		}
	}

	if !found {
		return SetRes{}, errors.New("Set not found")
	}
	log.Print(foundIdxS)

	err := set.EditSet(req.ActualRepCount)
	if err != nil {
		return SetRes{}, err
	}

	exercise.Sets[foundIdxS] = set
	t.wlog.Exercises[foundIdxE] = exercise

	res, err := t.repo.UpdateSet(req)
	if err != nil {
		log.Printf("%s: couldn't update Set from repo: %s", "tracker", err.Error())
		return SetRes{}, errInternal
	}

	return res, nil
}

func (t *tracker) FetchAllWorkoutLogs() ([]WorkoutLogRes, error) {
	var wlogsRes []WorkoutLogRes
	wlogsRes, err := t.repo.FindAllWorkoutLogsForAthlete(*t.ath)
	if err != nil {
		log.Printf("%s: couldn't fetch all WorkoutLogs from repo: %s", "tracker", err.Error())
		return wlogsRes, errInternal
	}

	return wlogsRes, nil
}

func (t *tracker) FetchAllExercisesForWorkoutLog() ([]ExerciseRes, error) {
	var exercisesRes []ExerciseRes

	if t.wlog == nil {
		return exercisesRes, errNilWorkoutLog
	}

	exercisesRes, err := t.repo.FindAllExercisesForWorkoutLog(*t.wlog)
	if err != nil {
		log.Printf("%s: couldn't fetch all Exercises from repo: %s", "tracker", err.Error())
		return exercisesRes, errInternal

	}

	return exercisesRes, nil
}

func (t *tracker) FetchAllSetsForExercise(exerciseID string) ([]SetRes, error) {
	var setsRes []SetRes

	if t.wlog == nil {
		return setsRes, errNilWorkoutLog
	}

	found := false
	var exercise wl.Exercise
	for _, e := range t.wlog.Exercises {
		if e.ExerciseID == exerciseID {
			found = true
			exercise = e
		}
	}

	if !found {
		return setsRes, errExerciseNotFound
	}

	setsRes, err := t.repo.FindAllSetsForExercise(exercise)
	if err != nil {
		log.Printf("%s: couldn't fetch all Sets from repo: %s", "tracker", err.Error())
		return setsRes, errInternal
	}

	return setsRes, nil
}
