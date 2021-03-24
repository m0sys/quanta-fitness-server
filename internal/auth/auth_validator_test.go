package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

func TestValidateRegisterationMismatch(t *testing.T) {

	testValidator := NewAuthValidator(nil)

	err := testValidator.ValidateRegisteration(MOCK_USERNAME, MOCK_EMAIL, MOCK_PWD, "hadio")

	assert.NotNil(t, err)
	assert.Equal(t, "Password must equal Confirm!", err.Error())
}

func TestValidateRegisterationUserExists(t *testing.T) {
	mockStore := new(MockStore)

	var id int64 = 1
	user := entity.User{
		ID: id,
		BaseUser: entity.BaseUser{
			Username: MOCK_USERNAME,
			Email:    MOCK_EMAIL,
			Password: MOCK_PWD,
		},
		Weight: MOCK_WEIGHT,
		Height: MOCK_HEIGHT,
		Gender: MOCK_GENDER,
	}

	mockStore.On("FindUserByUsername").Return(user, true, nil)

	testValidator := NewAuthValidator(mockStore)

	err := testValidator.ValidateRegisteration(MOCK_USERNAME, MOCK_EMAIL, MOCK_PWD, MOCK_PWD)

	assert.NotNil(t, err)
	assert.Equal(t, "User already exists!", err.Error())
}

func TestValidateRegisterationSuccess(t *testing.T) {
	mockStore := new(MockStore)

	var id int64 = 1
	user := entity.User{
		ID: id,
		BaseUser: entity.BaseUser{
			Username: MOCK_USERNAME,
			Email:    MOCK_EMAIL,
			Password: MOCK_PWD,
		},
		Weight: MOCK_WEIGHT,
		Height: MOCK_HEIGHT,
		Gender: MOCK_GENDER,
	}

	mockStore.On("FindUserByUsername").Return(user, false, nil)

	testValidator := NewAuthValidator(mockStore)

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

	testValidator := NewAuthValidator(mockStore)

	err := testValidator.ValidateRegisteration(MOCK_USERNAME, notEmail, MOCK_PWD, MOCK_PWD)

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid email!", err.Error())
}
