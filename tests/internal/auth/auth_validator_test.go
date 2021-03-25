package authtests

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/mhd53/quanta-fitness-server/internal/auth"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

func TestValidateRegisterationMismatch(t *testing.T) {

	testValidator := auth.NewAuthValidator(nil)

	err := testValidator.ValidateRegisteration(MOCK_USERNAME, MOCK_EMAIL, MOCK_PWD, "hadio")

	assert.NotNil(t, err)
	assert.Equal(t, "Password must equal Confirm!", err.Error())
}

func TestValidateRegisterationUserExists(t *testing.T) {
	mockStore := new(MockStore)

	var id int64 = 1
	user := CreateValidMockUser(id)

	mockStore.On("FindUserByUsername").Return(user, true, nil)

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateRegisteration(MOCK_USERNAME, MOCK_EMAIL, MOCK_PWD, MOCK_PWD)

	assert.NotNil(t, err)
	assert.Equal(t, "User already exists!", err.Error())
}

func TestValidateRegisterationSuccess(t *testing.T) {
	mockStore := new(MockStore)

	mockStore.On("FindUserByUsername").Return(entity.User{}, false, nil)

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateRegisteration(MOCK_USERNAME, MOCK_EMAIL, MOCK_PWD, MOCK_PWD)

	assert.Nil(t, err)
}

func TestValidateRegisterationWithInvalidEmail(t *testing.T) {
	mockStore := new(MockStore)
	notEmail := "notanemail.com"

	var id int64 = 1
	user := entity.User{
		ID: id,
		BaseUser: entity.BaseUser{
			Username: MOCK_USERNAME,
			Email:    notEmail,
			Password: MOCK_PWD,
		},
		Weight: MOCK_WEIGHT,
		Height: MOCK_HEIGHT,
		Gender: MOCK_GENDER,
	}

	mockStore.On("FindUserByUsername").Return(user, false, nil)

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateRegisteration(MOCK_USERNAME, notEmail, MOCK_PWD, MOCK_PWD)

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid email!", err.Error())
}

func TestValidateLoginWithUnameWhenUserNotExist(t *testing.T) {
	mockStore := new(MockStore)

	mockStore.On("FindUserByUsername").Return(entity.User{}, false, nil)

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateLoginWithUname(MOCK_USERNAME, MOCK_PWD)

	assert.NotNil(t, err)
	assert.Equal(t, "Username doesn't exist!", err.Error())
}

func TestValidateLoginWithUnameWhenIncorrectPwd(t *testing.T) {
	mockStore := new(MockStore)

	var id int64 = 1
	user := CreateValidMockUser(id)

	mockStore.On("FindUserByUsername").Return(user, true, nil)

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateLoginWithUname(MOCK_USERNAME, "bobin")

	assert.NotNil(t, err)
	assert.Equal(t, "Incorrect password!", err.Error())
}

func TestValidateLoginWithUnameSuccess(t *testing.T) {
	mockStore := new(MockStore)

	var id int64 = 1
	user := CreateValidMockUser(id)
	mockStore.On("FindUserByUsername").Return(user, true, nil)

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateLoginWithUname(MOCK_USERNAME, MOCK_PWD)

	assert.Nil(t, err)
}

func TestValidateLoginWithEmaileWhenUserNotExist(t *testing.T) {
	mockStore := new(MockStore)

	mockStore.On("FindUserByEmail").Return(entity.User{}, false, nil)

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateLoginWithEmail(MOCK_EMAIL, MOCK_PWD)

	assert.NotNil(t, err)
	assert.Equal(t, "Email doesn't exist!", err.Error())
}

func TestValidateLoginWithEmailWhenIncorrectPwd(t *testing.T) {
	mockStore := new(MockStore)

	var id int64 = 1
	user := CreateValidMockUser(id)

	mockStore.On("FindUserByEmail").Return(user, true, nil)

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateLoginWithEmail(MOCK_EMAIL, "bobin")

	assert.NotNil(t, err)
	assert.Equal(t, "Incorrect password!", err.Error())
}

func TestValidateLoginWithEmailWhenInvalidEmail(t *testing.T) {
	mockStore := new(MockStore)

	var id int64 = 1
	user := CreateValidMockUser(id)

	mockStore.On("FindUserByEmail").Return(user, true, nil)

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateLoginWithEmail("notanemail.com", MOCK_PWD)

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid email!", err.Error())
}

func TestValidateLoginWithEmailSuccess(t *testing.T) {
	mockStore := new(MockStore)

	var id int64 = 1
	user := CreateValidMockUser(id)
	mockStore.On("FindUserByEmail").Return(user, true, nil)

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateLoginWithEmail(MOCK_EMAIL, MOCK_PWD)

	assert.Nil(t, err)
}
