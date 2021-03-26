package workoutstore

import (
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type WorkoutStore interface {
	Save(workout entity.BaseWorkout) (entity.Workout, error)
	Update(wid int64, updates entity.BaseWorkout) error
	FindWorkoutById(wid int64) (entity.Workout, bool, error)
	DeleteWorkout(wid int64) error
	FindAllWorkoutsByUname(uname string) ([]entity.Workout, error)
}
