package trackertest

import (
	"os"
	"testing"

	"github.com/mhd53/quanta-fitness-server/athlete"
	"github.com/mhd53/quanta-fitness-server/internal/repository/inmem/trackerrepo"
	tckr "github.com/mhd53/quanta-fitness-server/tracker"
)

var (
	testAthlete athlete.Athlete
	testTracker tckr.WorkoutTracker
)

func TestMain(m *testing.M) {
	testRepo := trackerrepo.NewTrackerRepo()
	testAthlete = athlete.NewAthlete()
	testTracker = tckr.NewTracker(testRepo, &testAthlete)
	os.Exit(m.Run())
}
