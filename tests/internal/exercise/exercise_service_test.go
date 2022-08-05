package exercisetest

import (
	"github.com/stretchr/testify/assert"
	"testing"

	es "github.com/m0sys/quanta-fitness-server/internal/datastore/exercisestore"
	us "github.com/m0sys/quanta-fitness-server/internal/datastore/userstore"
	ws "github.com/m0sys/quanta-fitness-server/internal/datastore/workoutstore"
	w "github.com/m0sys/quanta-fitness-server/internal/workout"

	"github.com/m0sys/quanta-fitness-server/internal/entity"
	e "github.com/m0sys/quanta-fitness-server/internal/exercise"
	ats "github.com/m0sys/quanta-fitness-server/tests/internal/auth"
	wts "github.com/m0sys/quanta-fitness-server/tests/internal/workout"
)

func TestAddExerciseToWorkoutWhenUnauthenticated(t *testing.T) {
	testService, _, _, _ := setupService()

	created, err := testService.AddExerciseToWorkout(MOCK_VALID_NAME, "bobin", 0)

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
	assert.Empty(t, created)

}

func TestAddExerciseToWorkoutWhenWorkoutNotOwned(t *testing.T) {
	testService, mockUS, _, mockWS := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	wcreated, _ := mockWS.Save(entity.BaseWorkout{
		Username: "bobin",
		Title:    wts.MOCK_TITLE,
	})
	assert.NotEmpty(t, wcreated)

	created, err := testService.AddExerciseToWorkout(MOCK_VALID_NAME, "robin", 0)

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
	assert.Empty(t, created)

}

func TestAddExerciseToWorkoutWhenInvalidName(t *testing.T) {
	testService, mockUS, _, mockWS := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	wcreated, _ := mockWS.Save(wts.CreateValidMockBaseWorkout())
	assert.NotEmpty(t, wcreated)

	created, err := testService.AddExerciseToWorkout(MOCK_INVALID_NAME, "robin", 0)

	assert.NotNil(t, err)
	assert.Equal(t, "Name must be less than 38 characters!", err.Error())
	assert.Empty(t, created)

}

func TestAddExerciseToWorkoutSuccesss(t *testing.T) {
	testService, mockUS, _, mockWS := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	wcreated, _ := mockWS.Save(wts.CreateValidMockBaseWorkout())
	assert.NotEmpty(t, wcreated)

	created, err := testService.AddExerciseToWorkout(MOCK_VALID_NAME, "robin", 0)

	assert.Nil(t, err)
	assert.NotEmpty(t, created)

	// TODO: Test by checking database.

}

func TestUpdateExerciseWhenUnauthenticated(t *testing.T) {
	testService, _, _, _ := setupService()

	err := testService.UpdateExercise(0, "bobin", CreateMockValidUpdateExercise())

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
}

func TestUpdateExerciseWhenInvalidUpdate(t *testing.T) {
	testService, mockUS, mockES, _ := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	created, _ := mockES.Save(CreateMockValidBaseExercise())
	assert.NotEmpty(t, created)

	err := testService.UpdateExercise(0, "robin", CreateMockInvalidUpdateExercise())

	assert.NotNil(t, err)
	assert.Equal(t, "Name must be less than 38 characters!", err.Error())
}

func TestUpdateExerciseWhenSuccess(t *testing.T) {
	testService, mockUS, mockES, _ := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	created, _ := mockES.Save(CreateMockValidBaseExercise())
	assert.NotEmpty(t, created)

	err := testService.UpdateExercise(0, "robin", CreateMockValidUpdateExercise())

	assert.Nil(t, err)

	// TODO: Test update fields.
}

func TestGetExerciseWhenUnauthenticated(t *testing.T) {
	testService, _, _, _ := setupService()

	got, err := testService.GetExercise(0, "robin")

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
	assert.Empty(t, got)

}

func TestGetExerciseWhenExerciseNotExist(t *testing.T) {
	testService, mockUS, _, _ := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	got, err := testService.GetExercise(0, "robin")

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
	assert.Empty(t, got)

}

func TestGetExerciseSuccess(t *testing.T) {
	testService, mockUS, mockES, _ := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	created, _ := mockES.Save(CreateMockValidBaseExercise())
	assert.NotEmpty(t, created)

	got, err := testService.GetExercise(0, "robin")

	assert.Nil(t, err)
	assert.NotEmpty(t, got)
	assert.Equal(t, int64(0), got.ID)
	assert.Equal(t, MOCK_VALID_NAME, got.Name)
}

func TestGetExerciseForWorkoutWhenUnauthenticated(t *testing.T) {
	testService, _, _, _ := setupService()

	got, err := testService.GetExercisesForWorkout(0, "robin")

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
	assert.Empty(t, got)
}

func TestGetExerciseForWorkoutWhenNoWorkoutFound(t *testing.T) {
	testService, mockUS, _, _ := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	got, err := testService.GetExercisesForWorkout(0, "robin")

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
	assert.Empty(t, got)
}

func TestGetExerciseForWorkoutWhenNoExercisesFound(t *testing.T) {
	testService, mockUS, _, mockWS := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	wcreated, _ := mockWS.Save(wts.CreateValidMockBaseWorkout())
	assert.NotEmpty(t, wcreated)

	got, err := testService.GetExercisesForWorkout(0, "robin")

	assert.Nil(t, err)
	assert.Empty(t, got)
}

func TestGetExerciseForWorkoutWhenExercisesExist(t *testing.T) {
	testService, mockUS, mockES, mockWS := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	wcreated, _ := mockWS.Save(wts.CreateValidMockBaseWorkout())
	assert.NotEmpty(t, wcreated)

	created, _ := mockES.Save(CreateMockValidBaseExercise())
	assert.NotEmpty(t, created)

	created2, _ := mockES.Save(CreateMockValidBaseExercise())
	assert.NotEmpty(t, created2)

	created3, _ := mockES.Save(CreateMockValidBaseExercise())
	assert.NotEmpty(t, created3)

	got, err := testService.GetExercisesForWorkout(0, "robin")

	assert.Nil(t, err)
	assert.NotEmpty(t, got)
	assert.Equal(t, 3, len(got))
}

func TestGetExerciseForUserWhenUnauthenticated(t *testing.T) {
	skipTest(t)
	testService, _, _, _ := setupService()

	got, err := testService.GetExercisesForUser("robin")

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
	assert.Empty(t, got)
}

func TestDeleteExerciseWhenUnauthenticated(t *testing.T) {
	testService, _, _, _ := setupService()

	err := testService.DeleteExercise(0, "robin")

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
}

func TestDeleteExerciseWhenExerciseNotExist(t *testing.T) {
	testService, mockUS, mockES, _ := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	created, _ := mockES.Save(CreateMockValidBaseExercise())
	assert.NotEmpty(t, created)

	err := testService.DeleteExercise(0, "robin")

	assert.Nil(t, err)

	exercise, found, _ := mockES.FindExerciseById(0)

	assert.False(t, found)
	assert.Empty(t, exercise)

}

func TestDeleteExerciseSuccess(t *testing.T) {
	testService, mockUS, _, _ := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	err := testService.DeleteExercise(0, "robin")

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
}

// Utility funcs.

func setupService() (e.ExerciseService, us.UserStore, es.ExerciseStore, ws.WorkoutStore) {
	mockUS := us.NewMockUserStore()
	mockES := es.NewMockExerciseStore()
	mockWS := ws.NewMockWorkoutStore()

	testWauthorizer := w.NewWorkoutAuthorizer(mockWS, mockUS)
	testAuthorizer := e.NewExerciseAuthorizer(mockES, mockUS, testWauthorizer)
	testValidator := e.NewExerciseValidator()

	return e.NewExerciseService(mockES, testAuthorizer, testValidator), mockUS, mockES, mockWS
}

func skipTest(t *testing.T) {
	t.Skip("Implement AuthorizedReadAccess first!.")
}
