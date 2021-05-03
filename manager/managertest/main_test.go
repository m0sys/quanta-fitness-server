package managertest

import (
	"github.com/mhd53/quanta-fitness-server/athlete"
	"github.com/mhd53/quanta-fitness-server/internal/repository/inmem/managerrepo"
	"github.com/mhd53/quanta-fitness-server/manager"
)

func setup() manager.AthleteManager {
	testRepo := managerrepo.NewMangerRepo()
	testAthlete := athlete.NewAthlete()
	return manager.NewManager(testRepo, &testAthlete)
}
