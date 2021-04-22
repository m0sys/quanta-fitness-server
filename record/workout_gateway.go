package record

import (
	"github.com/mhd53/quanta-fitness-server/internal/id"
	w "github.com/mhd53/quanta-fitness-server/workout"
)

type WorkoutGateway interface {
	SaveWorkout(w.Workout) error
	FindWorkoutByID(id.ID) (w.Workout, error)
}
