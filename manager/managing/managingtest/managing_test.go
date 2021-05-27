package managiningtest

import (
	"testing"

	"github.com/mhd53/quanta-fitness-server/account/user"
	"github.com/mhd53/quanta-fitness-server/internal/random"
	m "github.com/mhd53/quanta-fitness-server/manager/managing"
	"github.com/mhd53/quanta-fitness-server/manager/managing/adapters"
	"github.com/stretchr/testify/require"
)

func TestCreateNewAthlete(t *testing.T) {
	service := setup()
	usr, err := user.NewUser(random.String(10), random.Email(), random.String(100))
	require.NoError(t, err)
	require.NotEmpty(t, usr)

	t.Run("When success", func(t *testing.T) {
		ath, err := service.CreateNewAthlete(usr)
		require.NoError(t, err)
		require.NotEmpty(t, ath)
	})
}

func TestFetchAthlete(t *testing.T) {
	service := setup()
	usr, err := user.NewUser(random.String(10), random.Email(), random.String(100))
	require.NoError(t, err)
	require.NotEmpty(t, usr)

	t.Run("When User not found", func(t *testing.T) {
		ath, err := service.FetchAthlete(usr)
		require.Error(t, err)
		require.Empty(t, ath)
		require.Equal(t, m.ErrAthleteNotFound.Error(), err.Error())
	})

	t.Run("When successs", func(t *testing.T) {
		ath, err := service.CreateNewAthlete(usr)
		require.NoError(t, err)
		require.NotEmpty(t, ath)

		athFetched, err := service.FetchAthlete(usr)
		require.NoError(t, err)
		require.NotEmpty(t, athFetched)
		require.Equal(t, ath.AthleteID(), athFetched.AthleteID())
	})
}

func setup() m.ManagingService {
	repo := adapters.NewInMemRepo()
	return m.NewManagingService(repo)
}
