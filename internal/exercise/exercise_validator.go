package exercise

import (
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type ExerciseValidator interface {
	ValidateCreateExercise(name string) error
	ValidateUpdateExercise(updates entity.ExerciseUpdate) error
}
