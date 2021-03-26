package workouttests

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"time"

	// "github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

var (
	MOCK_ERROR = errors.New("Mock Error")
)

const (
	MOCK_UNAME       = "robin"
	MOCK_TITLE       = "Chest Day"
	MOCK_VAL_TITLE   = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus."
	MOCK_INVAL_TITLE = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus.."
)

type MockStore struct {
	mock.Mock
}

func (mock *MockStore) Save(workout entity.BaseWorkout) (entity.Workout, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entity.Workout), args.Error(1)
}

func (mock *MockStore) Update(wid int64, updates entity.BaseWorkout) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockStore) FindWorkoutById(wid int64) (entity.Workout, bool, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entity.Workout), args.Bool(1), args.Error(2)
}

func (mock *MockStore) DeleteWorkout(wid int64) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *MockStore) FindAllWorkoutsByUname(uname string) ([]entity.Workout, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Workout), args.Error(1)
}

func CreateValidMockWorkout(id int64) entity.Workout {

	workout := entity.Workout{
		BaseWorkout: entity.BaseWorkout{
			Title:    MOCK_TITLE,
			Username: MOCK_UNAME,
		},
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return workout
}

func CreateValidMockBaseWorkout() entity.BaseWorkout {
	return entity.BaseWorkout{
		Title:    MOCK_VAL_TITLE,
		Username: MOCK_UNAME,
	}
}

func CreateInvalidMockBaseWorkout() entity.BaseWorkout {
	return entity.BaseWorkout{
		Title:    MOCK_INVAL_TITLE,
		Username: MOCK_UNAME,
	}
}
