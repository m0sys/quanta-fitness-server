package exercisetest

import (
	"github.com/stretchr/testify/assert"
	"testing"

	es "github.com/mhd53/quanta-fitness-server/internal/datastore/exercisestore"
	us "github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
	ws "github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"
	w "github.com/mhd53/quanta-fitness-server/internal/workout"

	"github.com/mhd53/quanta-fitness-server/internal/entity"
	e "github.com/mhd53/quanta-fitness-server/internal/exercise"
	ats "github.com/mhd53/quanta-fitness-server/tests/internal/auth"
	wts "github.com/mhd53/quanta-fitness-server/tests/internal/workout"
)

func TestAddExerciseToWorkoutWhenUnauthenticated(t *testing.T) {
	mockUS := us.NewMockUserStore()
	mockES := es.NewMockExerciseStore()
	mockWS := ws.NewMockWorkoutStore()

	testWauthorizer := w.NewWorkoutAuthorizer(mockWS, mockUS)
	testAuthorizer := e.NewExerciseAuthorizer(mockES, mockUS, testWauthorizer)
	testValidator := e.NewExerciseValidator()

	testService := e.NewExerciseService(mockES, testAuthorizer, testValidator)

	created, err := testService.AddExerciseToWorkout(MOCK_VALID_NAME, "bobin", 0)

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
	assert.Empty(t, created)

}

func TestAddExerciseToWorkoutWhenWorkoutNotOwned(t *testing.T) {
	mockUS := us.NewMockUserStore()
	mockES := es.NewMockExerciseStore()
	mockWS := ws.NewMockWorkoutStore()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	wcreated, _ := mockWS.Save(entity.BaseWorkout{
		Username: "bobin",
		Title:    wts.MOCK_TITLE,
	})
	assert.NotEmpty(t, wcreated)

	testWauthorizer := w.NewWorkoutAuthorizer(mockWS, mockUS)
	testAuthorizer := e.NewExerciseAuthorizer(mockES, mockUS, testWauthorizer)
	testValidator := e.NewExerciseValidator()

	testService := e.NewExerciseService(mockES, testAuthorizer, testValidator)

	created, err := testService.AddExerciseToWorkout(MOCK_VALID_NAME, "robin", 0)

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
	assert.Empty(t, created)

}

func TestAddExerciseToWorkoutWhenInvalidName(t *testing.T) {
	mockUS := us.NewMockUserStore()
	mockES := es.NewMockExerciseStore()
	mockWS := ws.NewMockWorkoutStore()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	wcreated, _ := mockWS.Save(wts.CreateValidMockBaseWorkout())
	assert.NotEmpty(t, wcreated)

	testWauthorizer := w.NewWorkoutAuthorizer(mockWS, mockUS)
	testAuthorizer := e.NewExerciseAuthorizer(mockES, mockUS, testWauthorizer)
	testValidator := e.NewExerciseValidator()

	testService := e.NewExerciseService(mockES, testAuthorizer, testValidator)

	created, err := testService.AddExerciseToWorkout(MOCK_INVALID_NAME, "robin", 0)

	assert.NotNil(t, err)
	assert.Equal(t, "Name must be less than 38 characters!", err.Error())
	assert.Empty(t, created)

}

func TestAddExerciseToWorkoutSuccesss(t *testing.T) {
	mockUS := us.NewMockUserStore()
	mockES := es.NewMockExerciseStore()
	mockWS := ws.NewMockWorkoutStore()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	wcreated, _ := mockWS.Save(wts.CreateValidMockBaseWorkout())
	assert.NotEmpty(t, wcreated)

	testWauthorizer := w.NewWorkoutAuthorizer(mockWS, mockUS)
	testAuthorizer := e.NewExerciseAuthorizer(mockES, mockUS, testWauthorizer)
	testValidator := e.NewExerciseValidator()

	testService := e.NewExerciseService(mockES, testAuthorizer, testValidator)

	created, err := testService.AddExerciseToWorkout(MOCK_VALID_NAME, "robin", 0)

	assert.Nil(t, err)
	assert.NotEmpty(t, created)

	// TODO: Test by checking database.

}
