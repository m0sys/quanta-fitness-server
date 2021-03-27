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

func TestAuthorizeWorkoutAccessWhenUnauthorized(t *testing.T) {
	mockUS := us.NewMockUserStore()
	mockES := es.NewMockExerciseStore()
	mockWS := ws.NewMockWorkoutStore()

	testWauthorizer := w.NewWorkoutAuthorizer(mockWS, mockUS)
	testAuthorizer := e.NewExerciseAuthorizer(mockES, mockUS, testWauthorizer)

	ok, err := testAuthorizer.AuthorizeWorkoutAccess("hello", 0)

	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestAuthorizeWorkoutAccessWhenWorkoutNotFound(t *testing.T) {
	mockUS := us.NewMockUserStore()
	mockES := es.NewMockExerciseStore()
	mockWS := ws.NewMockWorkoutStore()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	testWauthorizer := w.NewWorkoutAuthorizer(mockWS, mockUS)
	testAuthorizer := e.NewExerciseAuthorizer(mockES, mockUS, testWauthorizer)

	ok, err := testAuthorizer.AuthorizeWorkoutAccess("robin", 0)

	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestAuthorizeWorkoutAccessWhenWorkoutNotOwn(t *testing.T) {
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

	ok, err := testAuthorizer.AuthorizeWorkoutAccess("robin", 0)

	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestAuthorizeWorkoutAccessSuccess(t *testing.T) {
	mockUS := us.NewMockUserStore()
	mockES := es.NewMockExerciseStore()
	mockWS := ws.NewMockWorkoutStore()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	wcreated, _ := mockWS.Save(wts.CreateValidMockBaseWorkout())
	assert.NotEmpty(t, wcreated)

	testWauthorizer := w.NewWorkoutAuthorizer(mockWS, mockUS)
	testAuthorizer := e.NewExerciseAuthorizer(mockES, mockUS, testWauthorizer)

	ok, err := testAuthorizer.AuthorizeWorkoutAccess("robin", 0)

	assert.Nil(t, err)
	assert.True(t, ok)
}
