package authtest

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/mhd53/quanta-fitness-server/api/auth"
)

func TestRegisterNewUserWhenUserExists(t *testing.T) {
	server := auth.NewServerAuth()

	token, err := server.RegisterNewUser("hero", "hero@gmail.com", "nero", "nero")
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	token2, err2 := server.RegisterNewUser("hero", "hero@gmail.com", "nero", "nero")
	assert.NotNil(t, err2)
	assert.Empty(t, token2)

}

func TestRegisterNewUserSuccess(t *testing.T) {
	server := auth.NewServerAuth()

	token, err := server.RegisterNewUser("hero", "hero@gmail.com", "nero", "nero")

	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}
