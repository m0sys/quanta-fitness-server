package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

func TestRegisterWhenUserExists(t *testing.T) {
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
	mockStore.On("Save").Return(user, nil)

	testValidator := NewAuthValidator(mockStore)
	testService := NewAuthService(mockStore, testValidator)

	token, err := testService.Register(MOCK_USERNAME, MOCK_EMAIL, MOCK_PWD, MOCK_PWD)

	assert.NotNil(t, err)
	assert.Empty(t, token)
}

func TestRegisterSuccess(t *testing.T) {
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

	mockStore.On("FindUserByUsername").Return(entity.User{}, false, nil)
	mockStore.On("Save").Return(user, nil)

	testValidator := NewAuthValidator(mockStore)
	testService := NewAuthService(mockStore, testValidator)

	token, err := testService.Register(MOCK_USERNAME, MOCK_EMAIL, MOCK_PWD, MOCK_PWD)

	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func TestLoginWithUname(t *testing.T) {
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

	mockStore.On("FindUserByUsername").Return(entity.User{}, false, nil)
	mockStore.On("Save").Return(user, nil)

	testValidator := NewAuthValidator(mockStore)
	testService := NewAuthService(mockStore, testValidator)

	token, err := testService.Register(MOCK_USERNAME, MOCK_EMAIL, MOCK_PWD, MOCK_PWD)

	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}
