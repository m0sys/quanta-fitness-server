package auth

import (
	"errors"
	"github.com/stretchr/testify/mock"

	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

var (
	MOCK_ERROR = errors.New("Mock Error")
)

const (
	MOCK_ID       int64   = 1
	MOCK_USERNAME         = "robin"
	MOCK_EMAIL            = "robin@gmail.com"
	MOCK_PWD              = "robinhood"
	MOCK_WEIGHT   float32 = 75.0
	MOCK_HEIGHT   float32 = 158.0
	MOCK_GENDER           = "male"
)

type MockStore struct {
	mock.Mock
}

func (mock *MockStore) Save(user entity.BaseUser) (entity.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entity.User), args.Error(1)
}

func (mock *MockStore) FindUserByUsername(username string) (entity.User, bool, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entity.User), args.Bool(1), args.Error(2)
}

func (mock *MockStore) FindUserByEmail(email string) (entity.User, bool, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entity.User), args.Bool(1), args.Error(2)
}
