package trackertest

import (
	"github.com/mhd53/quanta-fitness-server/athlete"
	"github.com/mhd53/quanta-fitness-server/internal/repository/inmem/trackerrepo"
	tckr "github.com/mhd53/quanta-fitness-server/tracker"
)

func setup() (tckr.WorkoutTracker, *athlete.Athlete) {
	testRepo := trackerrepo.NewTrackerRepo()
	testAthlete := athlete.NewAthlete()
	testTracker := tckr.NewTracker(testRepo, &testAthlete)
	return testTracker, &testAthlete
}
