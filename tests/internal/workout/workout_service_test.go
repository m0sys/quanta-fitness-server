package workouttests

import (
	"github.com/stretchr/testify/assert"
	"testing"

	us "github.com/m0sys/quanta-fitness-server/internal/datastore/userstore"
	ws "github.com/m0sys/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/m0sys/quanta-fitness-server/internal/entity"
	w "github.com/m0sys/quanta-fitness-server/internal/workout"
	ats "github.com/m0sys/quanta-fitness-server/tests/internal/auth"
)

func TestCreateWorkoutWhenUnauthorized(t *testing.T) {
	testService, _, _ := setupService()

	created, err := testService.CreateWorkout("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus.", "robin")

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
	assert.Empty(t, created)
}

func TestCreateWorkoutWhenTitleIsInvalid(t *testing.T) {
	testService, mockUS, _ := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	created, err := testService.CreateWorkout("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus..", "robin")

	assert.NotNil(t, err)
	assert.Equal(t, "Title must be less than 76 characters!", err.Error())
	assert.Empty(t, created)
}

func TestCreateWorkoutSuccess(t *testing.T) {
	testService, mockUS, _ := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	title := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus."
	uname := "robin"

	created, err := testService.CreateWorkout(title, uname)

	assert.Nil(t, err)
	assert.NotEmpty(t, created)
	assert.Equal(t, created.Username, uname)
	assert.Equal(t, created.Title, title)
}

func TestUpdateWorkoutWhenUnauthorized(t *testing.T) {
	testService, _, mockWS := setupService()

	mockBaseWorkout := CreateValidMockBaseWorkout()
	created, _ := mockWS.Save(mockBaseWorkout)
	assert.NotEmpty(t, created)

	err := testService.UpdateWorkout(0, mockBaseWorkout, "robin")

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
}

func TestUpdateWorkoutInvalidWorkout(t *testing.T) {
	testService, mockUS, mockWS := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	mockWorkout := CreateValidMockBaseWorkout()
	created, _ := mockWS.Save(mockWorkout)
	assert.NotEmpty(t, created)
	mockBaseWorkout := CreateInvalidMockBaseWorkout()

	err := testService.UpdateWorkout(0, mockBaseWorkout, "robin")

	assert.NotNil(t, err)
}

func TestUpdateWorkoutSuccess(t *testing.T) {
	testService, mockUS, mockWS := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	mockWorkout := CreateValidMockBaseWorkout()
	created, _ := mockWS.Save(mockWorkout)
	title := "Chest DAY!!"
	mockBaseWorkout := entity.BaseWorkout{
		Username: MOCK_UNAME,
		Title:    title,
	}

	err := testService.UpdateWorkout(0, mockBaseWorkout, "robin")

	assert.Nil(t, err)

	got, err2 := testService.GetWorkout(0, "robin")
	assert.Nil(t, err2)
	assert.NotEmpty(t, got)
	assert.Equal(t, title, got.Title)
	assert.Equal(t, "robin", got.Username)
	assert.Equal(t, created.ID, got.ID)
	assert.Equal(t, created.CreatedAt, got.CreatedAt)
	assert.NotEqual(t, created.UpdatedAt, got.UpdatedAt)
}

func TestGetWorkoutWhenUnauthorized(t *testing.T) {
	testService, mockUS, _ := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	title := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus."
	uname := "robin"

	created, _ := testService.CreateWorkout(title, uname)
	assert.NotEmpty(t, created)

	got, err := testService.GetWorkout(0, "bobin")

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
	assert.Empty(t, got)
}

func TestGetWorkoutWhenNotExist(t *testing.T) {
	testService, mockUS, _ := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	title := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus."
	uname := "robin"

	created, _ := testService.CreateWorkout(title, uname)
	assert.NotEmpty(t, created)

	got, err := testService.GetWorkout(1, uname)

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
	assert.Empty(t, got)
}

func TestGetWorkoutSuccess(t *testing.T) {
	testService, mockUS, _ := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	title := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus."
	uname := "robin"
	var id2 int64 = 0

	created, _ := testService.CreateWorkout(title, uname)
	assert.NotEmpty(t, created)

	got, err := testService.GetWorkout(id2, uname)

	assert.Nil(t, err)
	assert.NotEmpty(t, got)
	assert.Equal(t, uname, got.Username)
	assert.Equal(t, title, got.Title)
	assert.Equal(t, id2, got.ID)
	assert.NotEmpty(t, got.CreatedAt)
}

func TestDeleteWorkoutWhenUnauthorized(t *testing.T) {
	testService, mockUS, _ := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	title := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus."
	uname := "robin"

	created, _ := testService.CreateWorkout(title, uname)
	assert.NotEmpty(t, created)

	err := testService.DeleteWorkout(0, "bobin")

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())

	got, _ := testService.GetWorkout(0, "robin")
	assert.NotEmpty(t, got)
	assert.Equal(t, int64(0), got.ID)
}

func TestDeleteWorkoutSuccess(t *testing.T) {
	testService, mockUS, _ := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	title := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus."
	uname := "robin"

	created, _ := testService.CreateWorkout(title, uname)
	assert.NotEmpty(t, created)

	err := testService.DeleteWorkout(0, "robin")

	assert.Nil(t, err)

	got, err2 := testService.GetWorkout(0, "robin")
	assert.NotEmpty(t, err2)
	assert.Equal(t, "Access Denied!", err2.Error())
	assert.Empty(t, got)
}

func TestGetWorkoutsForUserWhenUnauthenticated(t *testing.T) {
	testService, _, _ := setupService()

	got, err := testService.GetWorkoutsForUser("robin")

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
	assert.Empty(t, got)
}

func TestGetWorkoutsForUserWhenNoWorkoutFound(t *testing.T) {
	testService, mockUS, _ := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	got, err := testService.GetWorkoutsForUser("robin")

	assert.Nil(t, err)
	assert.Empty(t, got)
}

func TestGetWorkoutseForUserWhenNoWorkoutsFound(t *testing.T) {
	testService, mockUS, _ := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	got, err := testService.GetWorkoutsForUser("robin")

	assert.Nil(t, err)
	assert.Empty(t, got)
}

func TestGetWorkoutsForUserWhenWorkoutsExist(t *testing.T) {
	testService, mockUS, _ := setupService()

	ucreated, _ := mockUS.Save(ats.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	title := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus."
	uname := "robin"

	created, _ := testService.CreateWorkout(title, uname)
	assert.NotEmpty(t, created)

	created2, _ := testService.CreateWorkout(title, uname)
	assert.NotEmpty(t, created2)

	created3, _ := testService.CreateWorkout(title, uname)
	assert.NotEmpty(t, created3)

	got, err := testService.GetWorkoutsForUser("robin")

	assert.Nil(t, err)
	assert.NotEmpty(t, got)
	assert.Equal(t, 3, len(got))
}

// Utility funcs.

func setupService() (w.WorkoutService, us.UserStore, ws.WorkoutStore) {
	mockUS := us.NewMockUserStore()
	mockWS := ws.NewMockWorkoutStore()

	testAuthorizer := w.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := w.NewWorkoutValidator()
	testService := w.NewWorkoutService(
		mockWS,
		testAuthorizer,
		testValidator,
	)

	return testService, mockUS, mockWS
}
