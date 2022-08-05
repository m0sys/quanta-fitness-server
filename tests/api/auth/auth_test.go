package authtest

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/m0sys/quanta-fitness-server/api/auth"
	ustore "github.com/m0sys/quanta-fitness-server/internal/datastore/userstore"
)

func TestRegisterNewUserWhenUserExists(t *testing.T) {
	mockUS := ustore.NewMockUserStore()
	server := auth.NewServerAuth(mockUS)

	token, err := server.RegisterNewUser("hero", "hero@gmail.com", "nero", "nero")
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	token2, err2 := server.RegisterNewUser("hero", "hero@gmail.com", "nero", "nero")
	assert.NotNil(t, err2)
	assert.Empty(t, token2)

}

func TestRegisterNewUserSuccess(t *testing.T) {
	mockUS := ustore.NewMockUserStore()
	server := auth.NewServerAuth(mockUS)

	token, err := server.RegisterNewUser("hero", "hero@gmail.com", "nero", "nero")

	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func TestLoginWithUnameWhenUserNotExists(t *testing.T) {
	mockUS := ustore.NewMockUserStore()
	server := auth.NewServerAuth(mockUS)

	token, err := server.LoginWithUname("hero", "nero")

	assert.NotNil(t, err)
	assert.Equal(t, "Username doesn't exist!", err.Error())
	assert.Empty(t, token)

}

func TestLoginWithUnameWhenUserExists(t *testing.T) {
	mockUS := ustore.NewMockUserStore()
	server := auth.NewServerAuth(mockUS)

	_, err := server.RegisterNewUser("hero", "hero@gmail.com", "nero", "nero")
	assert.Nil(t, err)

	token, err2 := server.LoginWithUname("hero", "nero")

	assert.Nil(t, err2)
	assert.NotEmpty(t, token)

}

func TestLoginWithEmailWhenUserNotExists(t *testing.T) {
	mockUS := ustore.NewMockUserStore()
	server := auth.NewServerAuth(mockUS)

	token, err := server.LoginWithEmail("hero@gmail.com", "nero")

	assert.NotNil(t, err)
	assert.Equal(t, "Email doesn't exist!", err.Error())
	assert.Empty(t, token)

}

func TestLoginWithEmailSuccess(t *testing.T) {
	mockUS := ustore.NewMockUserStore()
	server := auth.NewServerAuth(mockUS)

	_, err := server.RegisterNewUser("hero", "hero@gmail.com", "nero", "nero")
	assert.Nil(t, err)

	token, err2 := server.LoginWithEmail("hero@gmail.com", "nero")

	assert.Nil(t, err2)
	assert.NotEmpty(t, token)
}
