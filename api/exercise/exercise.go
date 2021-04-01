package exercise

import (
	"fmt"

	estore "github.com/mhd53/quanta-fitness-server/internal/datastore/exercisestore"
	ustore "github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
	wstore "github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
	serv "github.com/mhd53/quanta-fitness-server/internal/exercise"
	wServ "github.com/mhd53/quanta-fitness-server/internal/workout"
	"github.com/mhd53/quanta-fitness-server/pkg/format"
)

var (
	service serv.ExerciseService
)

type server struct{}

type ExerciseServer interface {
	AddExerciseToWorkout(name, uname, wid string) (entity.Exercise, error)
	UpdateExercise(id, uname string, updates entity.ExerciseUpdate) (bool, error)
	GetExercise(id, uname string) (entity.Exercise, error)
	GetExercisesForWorkout(wid, uname string) ([]entity.Exercise, error)
	DeleteExercise(id, uname string) (bool, error)
}

func NewExerciseServer(us ustore.UserStore, ws wstore.WorkoutStore, es estore.ExerciseStore) ExerciseServer {
	wAuthorizer := wServ.NewWorkoutAuthorizer(ws, us)
	authorizer := serv.NewExerciseAuthorizer(es, us, wAuthorizer)
	validator := serv.NewExerciseValidator()
	service = serv.NewExerciseService(es, authorizer, validator)
	return &server{}
}

func (*server) AddExerciseToWorkout(name, uname, wid string) (entity.Exercise, error) {

	intID, err := format.ConvertToBase64(wid)
	if err != nil {
		return entity.Exercise{}, formatErr(err)

	}

	return service.AddExerciseToWorkout(name, uname, intID)
}

func (*server) UpdateExercise(id, uname string, updates entity.ExerciseUpdate) (bool, error) {

	intID, err := format.ConvertToBase64(id)
	if err != nil {
		return false, formatErr(err)
	}

	err2 := service.UpdateExercise(intID, uname, updates)

	if err2 != nil {
		return false, err2
	}

	return true, nil
}

func (*server) GetExercise(id, uname string) (entity.Exercise, error) {

	intID, err := format.ConvertToBase64(id)
	if err != nil {
		return entity.Exercise{}, formatErr(err)

	}

	got, err2 := service.GetExercise(intID, uname)

	if err2 != nil {
		return entity.Exercise{}, err2
	}

	return got, nil

}

func (*server) GetExercisesForWorkout(wid, uname string) ([]entity.Exercise, error) {

	intID, err := format.ConvertToBase64(wid)
	if err != nil {
		return []entity.Exercise{}, formatErr(err)

	}

	got, err2 := service.GetExercisesForWorkout(intID, uname)
	if err2 != nil {
		return []entity.Exercise{}, err2
	}

	return got, nil

}

func (*server) DeleteExercise(id, uname string) (bool, error) {

	intID, err := format.ConvertToBase64(id)
	if err != nil {
		return false, formatErr(err)

	}

	err2 := service.DeleteExercise(intID, uname)
	if err2 != nil {
		return false, err2
	}

	return true, nil

}

func formatErr(err error) error {
	return fmt.Errorf("%s: couldn't format id: %w", "API Exercise", err)
}
