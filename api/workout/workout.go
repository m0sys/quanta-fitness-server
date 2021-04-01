package workout

import (
	"fmt"

	ustore "github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
	wstore "github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
	serv "github.com/mhd53/quanta-fitness-server/internal/workout"
	"github.com/mhd53/quanta-fitness-server/pkg/format"
)

var (
	service serv.WorkoutService
)

type server struct{}

type WorkoutServer interface {
	CreateWorkout(title, uname string) (entity.Workout, error)
	UpdateWorkout(wid string, workout entity.BaseWorkout, uname string) (bool, error)
	DeleteWorkout(wid string, uname string) (bool, error)
	GetWorkout(wid string, uname string) (entity.Workout, error)
	GetWorkouts(uname string) ([]entity.Workout, error)
}

func NewWorkoutServer(us ustore.UserStore, ws wstore.WorkoutStore) WorkoutServer {
	authorizer := serv.NewWorkoutAuthorizer(ws, us)
	validator := serv.NewWorkoutValidator()
	service = serv.NewWorkoutService(ws, authorizer, validator)
	return &server{}
}
func (*server) CreateWorkout(title, uname string) (entity.Workout, error) {
	return service.CreateWorkout(title, uname)
}

func (*server) UpdateWorkout(wid string, workout entity.BaseWorkout, uname string) (bool, error) {

	intID, err := format.ConvertToBase64(wid)
	if err != nil {
		return false, formatErr(err)

	}

	err2 := service.UpdateWorkout(intID, workout, uname)
	if err2 != nil {
		return false, err2
	}

	return true, nil
}

func (*server) DeleteWorkout(wid string, uname string) (bool, error) {
	intID, err := format.ConvertToBase64(wid)
	if err != nil {
		return false, formatErr(err)

	}

	err2 := service.DeleteWorkout(intID, uname)
	if err2 != nil {
		return false, err2
	}

	return true, nil
}

func (*server) GetWorkout(wid string, uname string) (entity.Workout, error) {
	intID, err := format.ConvertToBase64(wid)
	if err != nil {
		return entity.Workout{}, formatErr(err)

	}

	return service.GetWorkout(intID, uname)
}

func (*server) GetWorkouts(uname string) ([]entity.Workout, error) {
	return service.GetWorkoutsForUser(uname)
}

func formatErr(err error) error {
	return fmt.Errorf("%s: couldn't format id: %w", "API Workout", err)
}
