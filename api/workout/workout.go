package workout

import (
	ustore "github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
	wstore "github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
	serv "github.com/mhd53/quanta-fitness-server/internal/workout"
	// "github.com/mhd53/quanta-fitness-server/pkg/crypto"
)

var (
	service serv.WorkoutService
)

type server struct{}

type WorkoutServer interface {
	CreateWorkout(title, uname string) (entity.Workout, error)
	UpdateWorkout(wid int64, workout entity.BaseWorkout, uname string) error
	DeleteWorkout(wid int64, uname string) error
	GetWorkout(wid int64, uname string) (entity.Workout, error)
}

func NewWorkoutServer(us ustore.UserStore, ws wstore.WorkoutStore) WorkoutServer {
	authorizer := serv.NewWorkoutAuthorizer(ws, us)
	validator := serv.NewWorkoutValidator()
	service = serv.NewWorkoutService(ws, authorizer, validator)
	return &server{}
}
func (*server) CreateWorkout(title, uname string) (entity.Workout, error) {
	created, err := service.CreateWorkout(title, uname)

	if err != nil {
		return entity.Workout{}, err
	}

	return created, nil
}

func (*server) UpdateWorkout(wid int64, workout entity.BaseWorkout, uname string) error {
	err := service.UpdateWorkout(wid, workout, uname)

	if err != nil {
		return err
	}

	return nil
}

func (*server) DeleteWorkout(wid int64, uname string) error {
	err := service.DeleteWorkout(wid, uname)
	if err != nil {
		return err
	}

	return nil

}

func (*server) GetWorkout(wid int64, uname string) (entity.Workout, error) {
	got, err := service.GetWorkout(wid, uname)
	if err != nil {
		return entity.Workout{}, err
	}

	return got, nil

}
