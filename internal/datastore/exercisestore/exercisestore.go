package exercisestore

import (
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type ExerciseStore interface {
	Save(exercise entity.BaseExercise) (entity.Exercise, error)
	Update(eid string, updates entity.ExerciseUpdate) error
	Delete(eid string) error
	FindExerciseById(eid string) (entity.Exercise, bool, error)
	FindAllExercisesByWID(wid string) ([]entity.Exercise, error)
	FindAllExercisesByUname(uname string) ([]entity.Exercise, error)
}
