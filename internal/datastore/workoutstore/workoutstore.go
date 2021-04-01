package workoutstore

import (
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type WorkoutStore interface {
	Save(workout entity.BaseWorkout) (entity.Workout, error)
	Update(wid string, updates entity.BaseWorkout) error
	FindWorkoutById(wid string) (entity.Workout, bool, error)
	DeleteWorkout(wid string) error
	FindAllWorkoutsByUname(uname string) ([]entity.Workout, error)
}
