package exercisestore

import (
	"errors"
	"time"

	"github.com/m0sys/quanta-fitness-server/internal/entity"
)

type store struct {
	exercises map[int64]entity.Exercise
	lastID    int64
}

func NewMockExerciseStore() ExerciseStore {
	return &store{exercises: make(map[int64]entity.Exercise)}
}

func (s *store) Save(exercise entity.BaseExercise) (entity.Exercise, error) {
	created := entity.Exercise{
		BaseExercise: exercise,
		ID:           s.lastID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Metrics: entity.Metrics{
			Weight:    0,
			TargetRep: 5,
			RestTime:  120,
			NumSets:   3,
		},
	}

	s.exercises[s.lastID] = created
	s.lastID += 1
	return created, nil
}

func (s *store) Update(eid int64, updates entity.ExerciseUpdate) error {
	prev := s.exercises[eid]
	updated := entity.Exercise{
		BaseExercise: entity.BaseExercise{
			Name:     updates.Name,
			WID:      prev.WID,
			Username: prev.Username,
		},
		CreatedAt: prev.CreatedAt,
		UpdatedAt: time.Now(),
		Metrics:   updates.Metrics,
	}

	s.exercises[eid] = updated
	return nil
}

func (s *store) Delete(eid int64) error {
	found := s.exercises[eid]
	if isEmpty(found) {
		return errors.New("Not Found!")
	}

	delete(s.exercises, eid)
	return nil
}

func isEmpty(found entity.Exercise) bool {
	return found == (entity.Exercise{})
}

func (s *store) FindExerciseById(eid int64) (entity.Exercise, bool, error) {
	found := s.exercises[eid]

	if isEmpty(found) {
		return entity.Exercise{}, false, nil
	}

	return found, true, nil
}

func (s *store) FindAllExercisesByWID(wid int64) ([]entity.Exercise, error) {
	var entries []entity.Exercise

	for k := range s.exercises {
		entry := s.exercises[k]
		if entry.WID == wid {
			entries = append(entries, entry)
		}
	}

	return entries, nil

}

func (s *store) FindAllExercisesByUname(uname string) ([]entity.Exercise, error) {
	var entries []entity.Exercise

	for k := range s.exercises {
		entry := s.exercises[k]
		if entry.Username == uname {
			entries = append(entries, entry)
		}
	}

	return entries, nil
}
