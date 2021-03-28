package esettest

import (
	"github.com/stretchr/testify/assert"
	"testing"

	// "github.com/mhd53/quanta-fitness-server/internal/entity"
	ess "github.com/mhd53/quanta-fitness-server/internal/datastore/esetstore"
	es "github.com/mhd53/quanta-fitness-server/internal/datastore/exercisestore"
	us "github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
	ws "github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"

	s "github.com/mhd53/quanta-fitness-server/internal/eset"
	e "github.com/mhd53/quanta-fitness-server/internal/exercise"
	w "github.com/mhd53/quanta-fitness-server/internal/workout"
	ats "github.com/mhd53/quanta-fitness-server/tests/internal/auth"
	ets "github.com/mhd53/quanta-fitness-server/tests/internal/exercise"
)

func TestAddEsetToExerciseWhenUnauthenticated(t *testing.T) {
	testService, _, _, _ := setupService()

	created, err := testService.AddEsetToExercise("robin", 0, 5, 120.1, 123.0)

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
	assert.Empty(t, created)
}

func TestAddEsetToExerciseExerciseNotExist(t *testing.T) {
	testService, mockUS, _, _ := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	created, err := testService.AddEsetToExercise("robin", 0, 5, 120.1, 123.0)

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
	assert.Empty(t, created)
}

func TestAddEsetToExerciseWhenInvalidEset(t *testing.T) {
	testService, mockUS, _, mockES := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	ecreated, _ := mockES.Save(ets.CreateMockValidBaseExercise())
	assert.NotEmpty(t, ecreated)

	created, err := testService.AddEsetToExercise("robin", 0, -5, 120.1, 123.0)

	assert.NotNil(t, err)
	assert.Equal(t, "Sign Error: Actual rep count must be positive or zero!", err.Error())
	assert.Empty(t, created)
}

func TestAddEsetToExerciseSuccess(t *testing.T) {
	testService, mockUS, _, mockES := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	ecreated, _ := mockES.Save(ets.CreateMockValidBaseExercise())
	assert.NotEmpty(t, ecreated)

	created, err := testService.AddEsetToExercise("robin", 0, 5, 120.1, 123.0)

	assert.Nil(t, err)
	assert.NotEmpty(t, created)
	assert.Equal(t, int64(0), created.ID)
}

// Utility funcs.

func setupService() (s.EsetService, us.UserStore, ess.EsetStore, es.ExerciseStore) {
	mockUS := us.NewMockUserStore()
	mockES := es.NewMockExerciseStore()
	mockWS := ws.NewMockWorkoutStore()
	mockESS := ess.NewMockEsetStore()

	testWauthorizer := w.NewWorkoutAuthorizer(mockWS, mockUS)
	testEAuthorizer := e.NewExerciseAuthorizer(mockES, mockUS, testWauthorizer)
	testAuthorizer := s.NewEsetAuthorizer(mockESS, mockUS, testEAuthorizer)
	testValidator := s.NewEsetValidator()

	return s.NewEsetService(mockESS, testAuthorizer, testValidator), mockUS, mockESS, mockES
}
