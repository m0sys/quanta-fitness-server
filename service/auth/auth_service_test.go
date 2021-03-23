package auth

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/mhd53/quanta-fitness-server/entity"
)

type MockStore struct {
	mock.Mock
}

const (
	MOCK_ID       int64   = 1
	MOCK_USERNAME         = "robin"
	MOCK_EMAIL            = "robin@gmail.com"
	MOCK_PASSWORD         = "robinhood"
	MOCK_WEIGHT   float32 = 75.0
	MOCK_HEIGHT   float32 = 158.0
	MOCK_GENDER           = "male"
)

func (mock *MockStore) Save(user *entity.UserRegister) (*entity.User, error) {
	args := mock.Called()
	result = args.Get(0)
	return result.(*entity.User), args.Error(1)
}

func (mock *MockStore) FindUserByUsername(username string) (*entity.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.User), args.Error(1)
}

func (mock *MockStore) FindUserByEmail(email string) (*entity.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.User), args.Error(1)
}

func TestValidateRegisterationMismatch(t *testing.T) {
	testService := NewAuthService(nil)
	user := entity.UserRegister{
		username: MOCK_USERNAME,
		email:    MOCK_EMAIL,
		password: MOCK_PASSWORD,
		confirm:  "bobin",
	}

	err := testService.ValidateRegisteration(&user)

	assert.NotNil(t, err)
	assert.Equal(t, "Password must equal Confirm!", err.Error())
}

/**
func TestRegister(t *testing.T) {
	mockStore := new(MockStore)
	user := entity.UserRegister{
		username:		MOCK_USERNAME,
		email:			MOCK_EMAIL,
		password:		MOCK_PASSWORD
		confirm: MOCK_PASSWORD
	}
}
*/
