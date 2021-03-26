package workouttests

import (
	"github.com/stretchr/testify/assert"
	"testing"

	ws "github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
	"github.com/mhd53/quanta-fitness-server/internal/workout"
	"github.com/mhd53/quanta-fitness-server/tests/internal/auth"
)

func TestCreateWorkoutWhenUnauthorized(t *testing.T) {
	mockUS := new(authtests.MockStore)
	mockWS := ws.NewMockWorkoutStore()
	mockUS.On("FindUserByUsername").Return(entity.User{}, false, nil)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator(mockWS)
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
	mockUS := new(authtests.MockStore)
	mockWS := ws.NewMockWorkoutStore()

	var id int64 = 1
	user := authtests.CreateValidMockUser(id)
	mockUS.On("FindUserByUsername").Return(user, true, nil)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator(mockWS)
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
	mockUS := new(authtests.MockStore)
	mockWS := ws.NewMockWorkoutStore()

	var id int64 = 1
	user := authtests.CreateValidMockUser(id)
	mockUS.On("FindUserByUsername").Return(user, true, nil)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator(mockWS)
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
	mockUS := new(authtests.MockStore)
	mockWS := ws.NewMockWorkoutStore()
	mockUS.On("FindUserByUsername").Return(entity.User{}, false, nil)

	mockBaseWorkout := CreateValidMockBaseWorkout()
	mockWS.Save(mockBaseWorkout)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator(mockWS)
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
	mockUS := new(authtests.MockStore)
	mockWS := ws.NewMockWorkoutStore()

	var id int64 = 1
	mockUser := authtests.CreateValidMockUser(id)
	mockUS.On("FindUserByUsername").Return(mockUser, true, nil)

	mockWorkout := CreateValidMockBaseWorkout()
	mockWS.Save(mockWorkout)
	mockBaseWorkout := CreateInvalidMockBaseWorkout()

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator(mockWS)
	testService := workout.NewWorkoutService(
		mockWS,
		testAuthorizer,
		testValidator,
	)

	err := testService.UpdateWorkout(0, mockBaseWorkout, "robin")

	assert.NotNil(t, err)
}

func TestUpdateWorkoutSuccess(t *testing.T) {
	mockUS := new(authtests.MockStore)
	mockWS := ws.NewMockWorkoutStore()

	var id int64 = 1
	mockUser := authtests.CreateValidMockUser(id)
	mockUS.On("FindUserByUsername").Return(mockUser, true, nil)

	mockWorkout := CreateValidMockBaseWorkout()
	created, _ := mockWS.Save(mockWorkout)
	title := "Chest DAY!!"
	mockBaseWorkout := entity.BaseWorkout{
		Username: MOCK_UNAME,
		Title:    title,
	}

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator(mockWS)
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
	mockUS := new(authtests.MockStore)
	mockWS := ws.NewMockWorkoutStore()
	var id int64 = 1
	user := authtests.CreateValidMockUser(id)
	mockUS.On("FindUserByUsername").Return(user, true, nil)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator(mockWS)
	testService := workout.NewWorkoutService(
		mockWS,
		testAuthorizer,
		testValidator,
	)

	title := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus."
	uname := "robin"

	testService.CreateWorkout(title, uname)

	got, err := testService.GetWorkout(0, "bobin")

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
	assert.Empty(t, got)
}

func TestGetWorkoutWhenNotExist(t *testing.T) {
	mockUS := new(authtests.MockStore)
	mockWS := ws.NewMockWorkoutStore()
	var id int64 = 1
	user := authtests.CreateValidMockUser(id)
	mockUS.On("FindUserByUsername").Return(user, true, nil)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator(mockWS)
	testService := workout.NewWorkoutService(
		mockWS,
		testAuthorizer,
		testValidator,
	)

	title := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus."
	uname := "robin"

	testService.CreateWorkout(title, uname)

	got, err := testService.GetWorkout(1, uname)

	assert.NotNil(t, err)
	assert.Equal(t, "Access Denied!", err.Error())
	assert.Empty(t, got)
}

func TestGetWorkoutSuccess(t *testing.T) {
	mockUS := new(authtests.MockStore)
	mockWS := ws.NewMockWorkoutStore()
	var id int64 = 1
	user := authtests.CreateValidMockUser(id)
	mockUS.On("FindUserByUsername").Return(user, true, nil)

	testAuthorizer := workout.NewWorkoutAuthorizer(mockWS, mockUS)
	testValidator := workout.NewWorkoutValidator(mockWS)
	testService := workout.NewWorkoutService(
		mockWS,
		testAuthorizer,
		testValidator,
	)

	title := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus."
	uname := "robin"
	var id2 int64 = 0

	testService.CreateWorkout(title, uname)

	got, err := testService.GetWorkout(id2, uname)

	assert.Nil(t, err)
	assert.NotEmpty(t, got)
	assert.Equal(t, uname, got.Username)
	assert.Equal(t, title, got.Title)
	assert.Equal(t, id2, got.ID)
	assert.NotEmpty(t, got.CreatedAt)
}
