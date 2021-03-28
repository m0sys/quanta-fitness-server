package esettest

import (
	"github.com/stretchr/testify/assert"
	"testing"

	esstore "github.com/mhd53/quanta-fitness-server/internal/datastore/esetstore"
	estore "github.com/mhd53/quanta-fitness-server/internal/datastore/exercisestore"
	ustore "github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
	wstore "github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
	esserv "github.com/mhd53/quanta-fitness-server/internal/eset"
	e "github.com/mhd53/quanta-fitness-server/internal/exercise"
	w "github.com/mhd53/quanta-fitness-server/internal/workout"
	ats "github.com/mhd53/quanta-fitness-server/tests/internal/auth"
	ets "github.com/mhd53/quanta-fitness-server/tests/internal/exercise"
)

func TestAuthorizeExerciseAccessWhenUnauthenticated(t *testing.T) {
	testAuth, _, _, _ := setupAuth()

	ok, err := testAuth.AuthorizeExerciseAccess("robin", 0)

	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestAuthorizeExerciseAccessWhenExerciseNotExist(t *testing.T) {
	testAuth, mockUS, _, _ := setupAuth()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	ok, err := testAuth.AuthorizeExerciseAccess("robin", 0)

	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestAuthorizeExerciseAccessWhenExerciseNotOwned(t *testing.T) {
	testAuth, mockUS, _, mockES := setupAuth()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	ecreated, _ := mockES.Save(entity.BaseExercise{
		Name:     ets.MOCK_VALID_NAME,
		WID:      0,
		Username: "bobin",
	})
	assert.NotEmpty(t, ecreated)

	ok, err := testAuth.AuthorizeExerciseAccess("robin", 0)

	assert.Nil(t, err)
	assert.False(t, ok)
}

func TestAuthorizeExerciseSuccess(t *testing.T) {
	testAuth, mockUS, _, mockES := setupAuth()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	ecreated, _ := mockES.Save(ets.CreateMockValidBaseExercise())
	assert.NotEmpty(t, ecreated)

	ok, err := testAuth.AuthorizeExerciseAccess("robin", 0)

	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestAuthorizeEsetAccessWhenUnauthenticated(t *testing.T) {
	testAuth, _, _, _ := setupAuth()

	ok, err := testAuth.AuthorizeEsetAccess("robin", 0)

	assert.Nil(t, err)
	assert.False(t, ok)

}

func TestAuthorizeEsetAccessWhenEsetNotExist(t *testing.T) {
	testAuth, mockUS, _, _ := setupAuth()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	ok, err := testAuth.AuthorizeEsetAccess("robin", 0)

	assert.Nil(t, err)
	assert.False(t, ok)

}

func TestAuthorizeEsetAccessWhenNotOwned(t *testing.T) {
	testAuth, mockUS, mockESS, _ := setupAuth()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	created, _ := mockESS.Save(CreateValidBaseBobinSet())
	assert.NotEmpty(t, created)

	ok, err := testAuth.AuthorizeEsetAccess("robin", 0)

	assert.Nil(t, err)
	assert.False(t, ok)

}

func TestAuthorizeEsetSuccess(t *testing.T) {
	testAuth, mockUS, mockESS, _ := setupAuth()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	created, _ := mockESS.Save(CreateValidBaseRobinSet())
	assert.NotEmpty(t, created)

	ok, err := testAuth.AuthorizeEsetAccess("robin", 0)

	assert.Nil(t, err)
	assert.True(t, ok)

}

// Util funcs.

func setupAuth() (esserv.EsetAuth, ustore.UserStore, esstore.EsetStore, estore.ExerciseStore) {
	mockUS := ustore.NewMockUserStore()
	mockESS := esstore.NewMockEsetStore()
	mockES := estore.NewMockExerciseStore()
	mockWS := wstore.NewMockWorkoutStore()

	testWauth := w.NewWorkoutAuthorizer(mockWS, mockUS)
	testEauth := e.NewExerciseAuthorizer(mockES, mockUS, testWauth)

	return esserv.NewEsetAuthorizer(mockESS, mockUS, testEauth), mockUS, mockESS, mockES
}
