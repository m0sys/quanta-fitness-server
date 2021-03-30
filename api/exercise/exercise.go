package exercise

import (
	"errors"
	"log"

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
		log.Panic("API Error: ", err.Error())
		return entity.Exercise{}, errors.New("Interna Error!")

	}

	return service.AddExerciseToWorkout(name, uname, intID)
}

func (*server) UpdateExercise(id, uname string, updates entity.ExerciseUpdate) (bool, error) {

	intID, err := format.ConvertToBase64(id)
	if err != nil {
		log.Panic("API Error: ", err.Error())
		return false, errors.New("Interna Error!")

	}

	err2 := service.UpdateExercise(intID, uname, updates)

	if err2 != nil {
		log.Panic("API Error: ", err2.Error())
		return false, errors.New("Interna Error!")
	}

	return true, nil
}

func (*server) GetExercise(id, uname string) (entity.Exercise, error) {

	intID, err := format.ConvertToBase64(id)
	if err != nil {
		log.Panic("API Error: ", err.Error())
		return entity.Exercise{}, errors.New("Interna Error!")

	}

	got, err2 := service.GetExercise(intID, uname)

	if err2 != nil {
		log.Panic("API Error: ", err2.Error())
		return entity.Exercise{}, errors.New("Interna Error!")
	}

	return got, nil

}

func (*server) GetExercisesForWorkout(wid, uname string) ([]entity.Exercise, error) {

	intID, err := format.ConvertToBase64(wid)
	if err != nil {
		log.Panic("API Error: ", err.Error())
		return []entity.Exercise{}, errors.New("Interna Error!")

	}

	got, err2 := service.GetExercisesForWorkout(intID, uname)
	if err2 != nil {
		log.Panic("API Error: ", err2.Error())
		return []entity.Exercise{}, errors.New("Interna Error!")
	}

	return got, nil

}

func (*server) DeleteExercise(id, uname string) (bool, error) {

	intID, err := format.ConvertToBase64(id)
	if err != nil {
		log.Panic("API Error: ", err.Error())
		return false, errors.New("Interna Error!")

	}

	err2 := service.DeleteExercise(intID, uname)
	if err2 != nil {
		log.Panic("API Error: ", err2.Error())
		return false, errors.New("Interna Error!")
	}

	return true, nil

}
