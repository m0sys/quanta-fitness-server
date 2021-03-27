package exercisestore

import (
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type ExerciseStore interface {
	Save(exercise entity.BaseExercise) (entity.Exercise, error)
	Update(eid int64, updates entity.ExerciseUpdate) error
	Delete(eid int64) error
	FindExerciseById(eid int64) (entity.Exercise, bool, error)
	FindAllExercisesByWID(wid int64) ([]entity.Exercise, error)
	FindAllExercisesByUname(uname string) ([]entity.Exercise, error)
}
