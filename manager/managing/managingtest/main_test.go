package managertest

import (
	"github.com/mhd53/quanta-fitness-server/internal/repository/inmem/managerrepo"
	"github.com/mhd53/quanta-fitness-server/manager"
	"github.com/mhd53/quanta-fitness-server/manager/athlete"
)

func setup() manager.AthleteManager {
	testRepo := managerrepo.NewMangerRepo()
	testAthlete := athlete.NewAthlete()
	return manager.NewManager(testRepo, &testAthlete)
}
