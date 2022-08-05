package workout

import (
	"errors"
	"fmt"
	"log"

	ws "github.com/m0sys/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/m0sys/quanta-fitness-server/internal/entity"
)

type WorkoutService interface {
	CreateWorkout(title string, uname string) (entity.Workout, error)
	UpdateWorkout(wid int64, workout entity.BaseWorkout, uname string) error
	DeleteWorkout(wid int64, uname string) error
	GetWorkout(wid int64, uname string) (entity.Workout, error)
	GetWorkoutsForUser(uname string) ([]entity.Workout, error)
	GetWorkoutExercises(wid int64, uname string) ([]entity.Exercise, error)
}

type service struct{}

var (
	sws         ws.WorkoutStore
	auth        WorkoutAuth
	val         WorkoutValidator
	deniedErr   = errors.New("Access Denied!")
	internalErr = errors.New("Internal Error!")
)

func NewWorkoutService(
	workoutstore ws.WorkoutStore,
	authorizer WorkoutAuth,
	validator WorkoutValidator) WorkoutService {
	sws = workoutstore
	auth = authorizer
	val = validator
	return &service{}
}

func (*service) CreateWorkout(title string, uname string) (entity.Workout, error) {
	authorized, err := auth.AuthorizeCreateWorkout(uname)

	if err != nil {
		logErr(err)
		return entity.Workout{}, internalErr
	}

	if !authorized {
		return entity.Workout{}, deniedErr

	}

	err2 := val.ValidateCreateWorkout(title)

	if err2 != nil {
		return entity.Workout{}, err2
	}
	workout := entity.BaseWorkout{
		Title:    title,
		Username: uname,
	}
	workoutDS, err3 := sws.Save(workout)

	if err3 != nil {
		return entity.Workout{}, formatErr(err3)
	}

	return workoutDS, nil
}

func (*service) UpdateWorkout(wid int64, workout entity.BaseWorkout, uname string) error {
	authorized, err := auth.AuthorizeAccessWorkout(uname, wid)

	if err != nil {
		return err
	}

	if !authorized {
		return deniedErr
	}

	err2 := val.ValidateUpdateWorkout(workout)

	if err2 != nil {
		return err2
	}

	err3 := sws.Update(wid, workout)

	if err3 != nil {
		return formatErr(err3)
	}

	return nil

}

func (*service) DeleteWorkout(wid int64, uname string) error {
	authorized, err := auth.AuthorizeAccessWorkout(uname, wid)
	if err != nil {
		return err
	}

	if !authorized {
		return deniedErr
	}

	err2 := sws.DeleteWorkout(wid)

	if err2 != nil {
		return formatErr(err2)
	}

	return nil

}

func (*service) GetWorkout(wid int64, uname string) (entity.Workout, error) {
	authorized, err := auth.AuthorizeAccessWorkout(uname, wid)
	if err != nil {
		logErr(err)
		return entity.Workout{}, internalErr
	}

	if !authorized {
		return entity.Workout{}, deniedErr
	}

	got, _, err2 := sws.FindWorkoutById(wid)

	if err2 != nil {
		return entity.Workout{}, formatErr(err2)
	}

	return got, nil
}

func (*service) GetWorkoutsForUser(uname string) ([]entity.Workout, error) {
	var workouts []entity.Workout

	authorized, err := auth.AuthorizeCreateWorkout(uname)
	if err != nil {
		logErr(err)
		return workouts, internalErr
	}

	if !authorized {
		return workouts, deniedErr
	}

	workouts, err2 := sws.FindAllWorkoutsByUname(uname)

	if err2 != nil {
		return workouts, formatErr(err2)
	}

	return workouts, nil
}

func (*service) GetWorkoutExercises(wid int64, uname string) ([]entity.Exercise, error) {
	var exercises []entity.Exercise
	return exercises, nil
}

func logErr(err error) {
	log.Printf("%s: %s", "workout_service", err.Error())
}

func formatErr(err error) error {
	return fmt.Errorf("%s: couldn't access db: %w", "workout_service", err)
}
