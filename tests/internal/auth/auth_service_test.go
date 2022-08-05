package authtests

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/m0sys/quanta-fitness-server/internal/auth"
	us "github.com/m0sys/quanta-fitness-server/internal/datastore/userstore"
	"github.com/m0sys/quanta-fitness-server/internal/entity"
	"github.com/m0sys/quanta-fitness-server/pkg/crypto"
)

func TestRegisterWhenUserExists(t *testing.T) {
	mockStore := us.NewMockUserStore()

	user := createAuthUser()
	created, _ := mockStore.Save(user)

	assert.NotEmpty(t, created)

	testValidator := auth.NewAuthValidator(mockStore)
	testService := auth.NewAuthService(mockStore, testValidator)

	userDS, err := testService.Register(MOCK_USERNAME, MOCK_EMAIL, MOCK_PWD, MOCK_PWD)

	assert.NotNil(t, err)
	assert.Equal(t, "Username already exists!", err.Error())
	assert.Empty(t, userDS)
}

func createAuthUser() entity.BaseUser {
	hashed, _ := crypto.HashPwd(MOCK_PWD)
	return entity.BaseUser{
		Username: MOCK_USERNAME,
		Email:    MOCK_EMAIL,
		Password: hashed,
	}
}

func createNewUser() entity.BaseUser {
	return entity.BaseUser{
		Username: MOCK_USERNAME,
		Email:    MOCK_EMAIL,
		Password: MOCK_PWD,
	}
}

func TestRegisterSuccess(t *testing.T) {
	mockStore := us.NewMockUserStore()

	user := createNewUser()

	testValidator := auth.NewAuthValidator(mockStore)
	testService := auth.NewAuthService(mockStore, testValidator)

	userDS, err := testService.Register(MOCK_USERNAME, MOCK_EMAIL, MOCK_PWD, MOCK_PWD)

	assert.Nil(t, err)
	assert.NotEmpty(t, user)
	assert.Equal(t, userDS.Username, MOCK_USERNAME)
}

func TestLoginWithUnameWhenUserNotExist(t *testing.T) {
	mockStore := us.NewMockUserStore()

	testValidator := auth.NewAuthValidator(mockStore)
	testService := auth.NewAuthService(mockStore, testValidator)

	userDS, err := testService.LoginWithUname(MOCK_USERNAME, MOCK_PWD)

	assert.NotNil(t, err)
	assert.Equal(t, "Username doesn't exist!", err.Error())
	assert.Empty(t, userDS)
}

func TestLoginWithUnameWhenIncorrectPwd(t *testing.T) {
	mockStore := us.NewMockUserStore()
	user := createAuthUser()
	created, _ := mockStore.Save(user)
	assert.NotEmpty(t, created)

	testValidator := auth.NewAuthValidator(mockStore)
	testService := auth.NewAuthService(mockStore, testValidator)

	userDS, err := testService.LoginWithUname(MOCK_USERNAME, "bobin")

	assert.NotNil(t, err)
	assert.Equal(t, "Incorrect password!", err.Error())
	assert.Empty(t, userDS)
}

func TestLoginWithUnameSuccess(t *testing.T) {
	mockStore := us.NewMockUserStore()
	user := createAuthUser()
	created, _ := mockStore.Save(user)
	assert.NotEmpty(t, created)

	testValidator := auth.NewAuthValidator(mockStore)
	testService := auth.NewAuthService(mockStore, testValidator)

	userDS, err := testService.LoginWithUname(MOCK_USERNAME, MOCK_PWD)

	assert.Nil(t, err)
	assert.NotEmpty(t, userDS)
	assert.Equal(t, userDS.Username, MOCK_USERNAME)
}

func TestLoginWithEmailWhenUserNotExist(t *testing.T) {
	mockStore := us.NewMockUserStore()

	testValidator := auth.NewAuthValidator(mockStore)
	testService := auth.NewAuthService(mockStore, testValidator)

	userDS, err := testService.LoginWithEmail(MOCK_EMAIL, MOCK_PWD)

	assert.NotNil(t, err)
	assert.Equal(t, "Email doesn't exist!", err.Error())
	assert.Empty(t, userDS)
}

func TestLoginWithEmailWhenIncorrectPwd(t *testing.T) {
	mockStore := us.NewMockUserStore()
	user := createAuthUser()
	created, _ := mockStore.Save(user)
	assert.NotEmpty(t, created)

	testValidator := auth.NewAuthValidator(mockStore)
	testService := auth.NewAuthService(mockStore, testValidator)

	userDS, err := testService.LoginWithEmail(MOCK_EMAIL, "bobin")

	assert.NotNil(t, err)
	assert.Equal(t, "Incorrect password!", err.Error())
	assert.Empty(t, userDS)
}

func TestLoginWithEmailWhenInvalidEmail(t *testing.T) {
	mockStore := us.NewMockUserStore()
	user := createAuthUser()
	created, _ := mockStore.Save(user)
	assert.NotEmpty(t, created)

	testValidator := auth.NewAuthValidator(mockStore)
	testService := auth.NewAuthService(mockStore, testValidator)

	userDS, err := testService.LoginWithEmail("notanemail", "bobin")

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid email address!", err.Error())
	assert.Empty(t, userDS)
}

func TestLoginWithEmailSuccess(t *testing.T) {
	mockStore := us.NewMockUserStore()
	user := createAuthUser()
	created, _ := mockStore.Save(user)
	assert.NotEmpty(t, created)

	testValidator := auth.NewAuthValidator(mockStore)
	testService := auth.NewAuthService(mockStore, testValidator)

	userDS, err := testService.LoginWithEmail(MOCK_EMAIL, MOCK_PWD)

	assert.Nil(t, err)
	assert.NotEmpty(t, userDS)
}
