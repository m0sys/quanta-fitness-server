package user

import (
	"testing"

	"github.com/mhd53/quanta-fitness-server/internal/random"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	t.Run("When username is less than 3 chars", func(t *testing.T) {
		user, err := NewUser(random.String(2), random.Email(), random.String(100))
		require.Error(t, err)
		require.Equal(t, errInvalidUname.Error(), err.Error())
		require.Empty(t, user)
	})

	t.Run("When password is less than 12 chars", func(t *testing.T) {
		user, err := NewUser(random.String(3), random.Email(), random.String(11))
		require.Error(t, err)
		require.Equal(t, errInvalidPassword.Error(), err.Error())
		require.Empty(t, user)
	})

	t.Run("When password is more than 128 chars", func(t *testing.T) {
		user, err := NewUser(random.String(3), random.Email(), random.String(129))
		require.Error(t, err)
		require.Equal(t, errInvalidPassword.Error(), err.Error())
		require.Empty(t, user)
	})

	t.Run("When email is not an email", func(t *testing.T) {
		user, err := NewUser(random.String(3), random.String(10), random.String(128))
		require.Error(t, err)
		require.Equal(t, errInvalidEmail.Error(), err.Error())
		require.Empty(t, user)
	})

	t.Run("When success", func(t *testing.T) {
		uname := random.String(3)
		email := random.Email()
		pwd := random.String(128)
		user, err := NewUser(uname, email, pwd)
		require.NoError(t, err)
		require.NotEmpty(t, user)
		require.Equal(t, uname, user.Username())
		require.Equal(t, email, user.Email())
		require.NotEqual(t, pwd, user.HashedPwd())
	})
}
