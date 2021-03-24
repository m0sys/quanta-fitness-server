package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

func TestRegisterWhenUserExists(t *testing.T) {
	mockStore := new(MockStore)

	var id int64 = 1
	user := CreateValidMockUser(id)

	mockStore.On("FindUserByUsername").Return(user, true, nil)
	mockStore.On("Save").Return(user, nil)

	testValidator := NewAuthValidator(mockStore)
	testService := NewAuthService(mockStore, testValidator)

	token, err := testService.Register(MOCK_USERNAME, MOCK_EMAIL, MOCK_PWD, MOCK_PWD)

	assert.NotNil(t, err)
	assert.Equal(t, "User already exists!", err.Error())
	assert.Empty(t, token)
}

func TestRegisterSuccess(t *testing.T) {
	mockStore := new(MockStore)

	var id int64 = 1
	user := CreateValidMockUser(id)

	mockStore.On("FindUserByUsername").Return(entity.User{}, false, nil)
	mockStore.On("Save").Return(user, nil)

	testValidator := NewAuthValidator(mockStore)
	testService := NewAuthService(mockStore, testValidator)

	token, err := testService.Register(MOCK_USERNAME, MOCK_EMAIL, MOCK_PWD, MOCK_PWD)

	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func TestLoginWithUnameWhenUserNotExist(t *testing.T) {
	mockStore := new(MockStore)

	mockStore.On("FindUserByUsername").Return(entity.User{}, false, nil)
	mockStore.On("Save").Return(entity.User{}, nil)

	testValidator := NewAuthValidator(mockStore)
	testService := NewAuthService(mockStore, testValidator)

	token, err := testService.LoginWithUname(MOCK_USERNAME, MOCK_PWD)

	assert.NotNil(t, err)
	assert.Equal(t, "Username doesn't exist!", err.Error())
	assert.Empty(t, token)
}

func TestLoginWithUnameWhenIncorrectPwd(t *testing.T) {
	mockStore := new(MockStore)

	var id int64 = 1
	user := CreateValidMockUser(id)

	mockStore.On("FindUserByUsername").Return(user, true, nil)
	mockStore.On("Save").Return(user, nil)

	testValidator := NewAuthValidator(mockStore)
	testService := NewAuthService(mockStore, testValidator)

	token, err := testService.LoginWithUname(MOCK_USERNAME, "bobin")

	assert.NotNil(t, err)
	assert.Equal(t, "Incorrect password!", err.Error())
	assert.Empty(t, token)
}

func TestLoginWithUnameSuccess(t *testing.T) {
	mockStore := new(MockStore)

	var id int64 = 1
	user := CreateValidMockUser(id)

	mockStore.On("FindUserByUsername").Return(user, true, nil)
	mockStore.On("Save").Return(user, nil)

	testValidator := NewAuthValidator(mockStore)
	testService := NewAuthService(mockStore, testValidator)

	token, err := testService.LoginWithUname(MOCK_USERNAME, MOCK_PWD)

	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func TestLoginWithEmailWhenUserNotExist(t *testing.T) {
	mockStore := new(MockStore)

	mockStore.On("FindUserByEmail").Return(entity.User{}, false, nil)
	mockStore.On("Save").Return(entity.User{}, nil)

	testValidator := NewAuthValidator(mockStore)
	testService := NewAuthService(mockStore, testValidator)

	token, err := testService.LoginWithEmail(MOCK_EMAIL, MOCK_PWD)

	assert.NotNil(t, err)
	assert.Equal(t, "Email doesn't exist!", err.Error())
	assert.Empty(t, token)
}

func TestLoginWithEmailWhenIncorrectPwd(t *testing.T) {
	mockStore := new(MockStore)

	var id int64 = 1
	user := CreateValidMockUser(id)

	mockStore.On("FindUserByEmail").Return(user, true, nil)
	mockStore.On("Save").Return(user, nil)

	testValidator := NewAuthValidator(mockStore)
	testService := NewAuthService(mockStore, testValidator)

	token, err := testService.LoginWithEmail(MOCK_EMAIL, "bobin")

	assert.NotNil(t, err)
	assert.Equal(t, "Incorrect password!", err.Error())
	assert.Empty(t, token)
}

func TestLoginWithEmailWhenInvalidEmail(t *testing.T) {
	mockStore := new(MockStore)

	var id int64 = 1
	user := CreateValidMockUser(id)

	mockStore.On("FindUserByEmail").Return(user, true, nil)
	mockStore.On("Save").Return(user, nil)

	testValidator := NewAuthValidator(mockStore)
	testService := NewAuthService(mockStore, testValidator)

	token, err := testService.LoginWithEmail("notanemail", "bobin")

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid email!", err.Error())
	assert.Empty(t, token)
}

func TestLoginWithEmailSuccess(t *testing.T) {
	mockStore := new(MockStore)

	var id int64 = 1
	user := CreateValidMockUser(id)

	mockStore.On("FindUserByEmail").Return(user, true, nil)
	mockStore.On("Save").Return(user, nil)

	testValidator := NewAuthValidator(mockStore)
	testService := NewAuthService(mockStore, testValidator)

	token, err := testService.LoginWithEmail(MOCK_EMAIL, MOCK_PWD)

	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}
