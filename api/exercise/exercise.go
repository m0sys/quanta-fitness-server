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
