package workouttests

import (
	"github.com/stretchr/testify/assert"
	"testing"

	us "github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
	ws "github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
	"github.com/mhd53/quanta-fitness-server/internal/workout"
	"github.com/mhd53/quanta-fitness-server/tests/internal/auth"
)

func TestCreateWorkoutWhenUnauthorized(t *testing.T) {
	mockUS := us.NewMockUserStore()
	mockWS := ws.NewMockWorkoutStore()

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator()
	testService := workout.NewWorkoutService(
		mockWS,
		testAuthorizer,
		testValidator,
	)

	created, err := testService.CreateWorkout("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus.", "robin")

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
	assert.Empty(t, created)
}

func TestCreateWorkoutWhenTitleIsInvalid(t *testing.T) {
	mockUS := us.NewMockUserStore()
	mockWS := ws.NewMockWorkoutStore()

	ucreated, _ := mockUS.Save(authtests.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator()
	testService := workout.NewWorkoutService(
		mockWS,
		testAuthorizer,
		testValidator,
	)

	created, err := testService.CreateWorkout("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus..", "robin")

	assert.NotNil(t, err)
	assert.Equal(t, "Title must be less than 76 characters!", err.Error())
	assert.Empty(t, created)
}

func TestCreateWorkoutSuccess(t *testing.T) {
	mockUS := us.NewMockUserStore()
	mockWS := ws.NewMockWorkoutStore()

	ucreated, _ := mockUS.Save(authtests.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator()
	testService := workout.NewWorkoutService(
		mockWS,
		testAuthorizer,
		testValidator,
	)

	title := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus."
	uname := "robin"

	created, err := testService.CreateWorkout(title, uname)

	assert.Nil(t, err)
	assert.NotEmpty(t, created)
	assert.Equal(t, created.Username, uname)
	assert.Equal(t, created.Title, title)
}

func TestUpdateWorkoutWhenUnauthorized(t *testing.T) {
	mockUS := us.NewMockUserStore()
	mockWS := ws.NewMockWorkoutStore()

	mockBaseWorkout := CreateValidMockBaseWorkout()
	created, _ := mockWS.Save(mockBaseWorkout)
	assert.NotEmpty(t, created)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator()
	testService := workout.NewWorkoutService(
		mockWS,
		testAuthorizer,
		testValidator,
	)

	err := testService.UpdateWorkout(0, mockBaseWorkout, "robin")

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
}

func TestUpdateWorkoutInvalidWorkout(t *testing.T) {
	mockUS := us.NewMockUserStore()
	mockWS := ws.NewMockWorkoutStore()

	ucreated, _ := mockUS.Save(authtests.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	mockWorkout := CreateValidMockBaseWorkout()
	created, _ := mockWS.Save(mockWorkout)
	assert.NotEmpty(t, created)
	mockBaseWorkout := CreateInvalidMockBaseWorkout()

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator()
	testService := workout.NewWorkoutService(
		mockWS,
		testAuthorizer,
		testValidator,
	)

	err := testService.UpdateWorkout(0, mockBaseWorkout, "robin")

	assert.NotNil(t, err)
}

func TestUpdateWorkoutSuccess(t *testing.T) {
	mockUS := us.NewMockUserStore()
	mockWS := ws.NewMockWorkoutStore()

	ucreated, _ := mockUS.Save(authtests.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	mockWorkout := CreateValidMockBaseWorkout()
	created, _ := mockWS.Save(mockWorkout)
	title := "Chest DAY!!"
	mockBaseWorkout := entity.BaseWorkout{
		Username: MOCK_UNAME,
		Title:    title,
	}

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator()
	testService := workout.NewWorkoutService(
		mockWS,
		testAuthorizer,
		testValidator,
	)

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
	mockUS := us.NewMockUserStore()
	mockWS := ws.NewMockWorkoutStore()

	ucreated, _ := mockUS.Save(authtests.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator()
	testService := workout.NewWorkoutService(
		mockWS,
		testAuthorizer,
		testValidator,
	)

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
	mockUS := us.NewMockUserStore()
	mockWS := ws.NewMockWorkoutStore()

	ucreated, _ := mockUS.Save(authtests.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator()
	testService := workout.NewWorkoutService(
		mockWS,
		testAuthorizer,
		testValidator,
	)

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
	mockUS := us.NewMockUserStore()
	mockWS := ws.NewMockWorkoutStore()

	ucreated, _ := mockUS.Save(authtests.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator()
	testService := workout.NewWorkoutService(
		mockWS,
		testAuthorizer,
		testValidator,
	)

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
	mockUS := us.NewMockUserStore()
	mockWS := ws.NewMockWorkoutStore()

	ucreated, _ := mockUS.Save(authtests.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator()
	testService := workout.NewWorkoutService(
		mockWS,
		testAuthorizer,
		testValidator,
	)

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
	mockUS := us.NewMockUserStore()
	mockWS := ws.NewMockWorkoutStore()

	ucreated, _ := mockUS.Save(authtests.CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator()
	testService := workout.NewWorkoutService(
		mockWS,
		testAuthorizer,
		testValidator,
	)

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
