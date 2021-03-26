package exercise

import (
	store "github.com/mhd53/quanta-fitness-server/internal/datastore/exercisestore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type ExerciseService interface {
	AddExerciseToWorkout(name, uname string, wid int64) (entity.Exercise, error)
	UpdateExercise(eid int64, uname string, updates entity.ExerciseUpdate) error
	DeleteExercise(eid int64, uname string) error
	GetExercise(eid int64, uname string) (entity.Exercise, error)
	GetExercisesForUser(uname string) ([]entity.Exercise, error)
}

type service struct{}

func NewExerciseService(estore store.ExerciseStore) ExerciseService {
	return &service{}
}

func (*service) AddExerciseToWorkout(name, uname string, wid int64) (entity.Exercise, error) {
	return entity.Exercise{}, nil

}

func (*service) UpdateExercise(eid int64, uname string, updates entity.ExerciseUpdate) error {
	return nil
}

func (*service) DeleteExercise(eid int64, uname string) error {
	return nil
}

func (*service) GetExercise(eid int64, uname string) (entity.Exercise, error) {
	return entity.Exercise{}, nil
}

func (*service) GetExercisesForUser(uname string) ([]entity.Exercise, error) {
	var exercises []entity.Exercise
	return exercises, nil
}
