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

func TestAuthorizeWorkoutAccessWhenUnauthenticated(t *testing.T) {
	testAuthorizer, _, _, _ := setupAuth()

	ok, err := testAuthorizer.AuthorizeWorkoutAccess("hello", 0)

	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestAuthorizeWorkoutAccessWhenWorkoutNotFound(t *testing.T) {
	testAuthorizer, mockUS, _, _ := setupAuth()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	ok, err := testAuthorizer.AuthorizeWorkoutAccess("robin", 0)

	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestAuthorizeWorkoutAccessWhenWorkoutNotOwn(t *testing.T) {
	testAuthorizer, mockUS, _, mockWS := setupAuth()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	wcreated, _ := mockWS.Save(entity.BaseWorkout{
		Username: "bobin",
		Title:    wts.MOCK_TITLE,
	})
	assert.NotEmpty(t, wcreated)

	ok, err := testAuthorizer.AuthorizeWorkoutAccess("robin", 0)

	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestAuthorizeWorkoutAccessSuccess(t *testing.T) {
	testAuthorizer, mockUS, _, mockWS := setupAuth()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	wcreated, _ := mockWS.Save(wts.CreateValidMockBaseWorkout())
	assert.NotEmpty(t, wcreated)

	ok, err := testAuthorizer.AuthorizeWorkoutAccess("robin", 0)

	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestAuthorizeExerciseAccessWhenUnauthenticated(t *testing.T) {
	testAuthorizer, _, _, _ := setupAuth()

	ok, err := testAuthorizer.AuthorizeExerciseAccess("hello", 0)

	assert.Nil(t, err)
	assert.False(t, ok)

}

func TestAuthorizeExerciseAccessWhenNotOwnExercise(t *testing.T) {
	testAuthorizer, mockUS, mockES, _ := setupAuth()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	created, _ := mockES.Save(entity.BaseExercise{
		Name:     MOCK_VALID_NAME,
		WID:      0,
		Username: "bobin",
	})
	assert.NotEmpty(t, created)

	ok, err := testAuthorizer.AuthorizeExerciseAccess("robin", 0)

	assert.Nil(t, err)
	assert.False(t, ok)

}

func TestAuthorizeExerciseAccessSuccess(t *testing.T) {
	testAuthorizer, mockUS, mockES, _ := setupAuth()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	created, _ := mockES.Save(CreateMockValidBaseExercise())
	assert.NotEmpty(t, created)

	ok, err := testAuthorizer.AuthorizeExerciseAccess("robin", 0)

	assert.Nil(t, err)
	assert.True(t, ok)

}

// Utility funcs.

func setupAuth() (e.ExerciseAuth, us.UserStore, es.ExerciseStore, ws.WorkoutStore) {
	mockUS := us.NewMockUserStore()
	mockES := es.NewMockExerciseStore()
	mockWS := ws.NewMockWorkoutStore()

	testWauthorizer := w.NewWorkoutAuthorizer(mockWS, mockUS)
	return e.NewExerciseAuthorizer(mockES, mockUS, testWauthorizer), mockUS, mockES, mockWS
}
