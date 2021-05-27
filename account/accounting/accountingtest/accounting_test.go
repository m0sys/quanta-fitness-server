package accountingtest

import (
	"testing"

	a "github.com/mhd53/quanta-fitness-server/account/accounting"
	"github.com/mhd53/quanta-fitness-server/account/accounting/adapters"
	"github.com/mhd53/quanta-fitness-server/internal/random"
	"github.com/stretchr/testify/require"
)

func TestSignUp(t *testing.T) {
	service := setup()

	t.Run("When success", func(t *testing.T) {
		user, err := service.SignUp(random.String(10), random.Email(), random.String(100))
		require.NoError(t, err)
		require.NotEmpty(t, user)
	})

	t.Run("When User with uname already taken", func(t *testing.T) {
		uname := random.String(10)
		user, err := service.SignUp(uname, random.Email(), random.String(100))
		require.NoError(t, err)
		require.NotEmpty(t, user)
		require.Equal(t, uname, user.Username())

		user, err = service.SignUp(uname, random.Email(), random.String(100))
		require.Error(t, err)
		require.Empty(t, user)
		require.Equal(t, a.ErrUnameAlreadyExists.Error(), err.Error())
	})

	t.Run("When User with email already taken", func(t *testing.T) {
		email := random.Email()
		user, err := service.SignUp(random.String(10), email, random.String(100))
		require.NoError(t, err)
		require.NotEmpty(t, user)
		require.Equal(t, email, user.Email())

		user, err = service.SignUp(random.String(10), email, random.String(100))
		require.Error(t, err)
		require.Empty(t, user)
		require.Equal(t, a.ErrEmailAlreadyExists.Error(), err.Error())
	})
}

func TestLogin(t *testing.T) {
	service := setup()
	t.Run("When User doesn't exist", func(t *testing.T) {
		user, err := service.Login(random.String(10), random.String(100))
		require.Error(t, err)
		require.Empty(t, user)
		require.Equal(t, a.ErrUnameNotFound.Error(), err.Error())
	})

	t.Run("When incorrect password", func(t *testing.T) {
		uname := random.String(10)
		password := random.String(100)

		user, err := service.SignUp(uname, random.Email(), password)
		require.NoError(t, err)
		require.NotEmpty(t, user)

		user, err = service.Login(uname, random.String(100))
		require.Error(t, err)
		require.Empty(t, user)
		require.Equal(t, a.ErrIncorrectPassword.Error(), err.Error())
	})

	t.Run("When success", func(t *testing.T) {
		uname := random.String(10)
		password := random.String(100)

		userSignUp, err := service.SignUp(uname, random.Email(), password)
		require.NoError(t, err)
		require.NotEmpty(t, userSignUp)

		userLogin, err := service.Login(uname, password)
		require.NoError(t, err)
		require.NotEmpty(t, userLogin)
		require.Equal(t, uname, userLogin.Username())
	})
}

func setup() a.AccountService {
	repo := adapters.NewInMemRepo()
	return a.NewAccountService(repo)

}
