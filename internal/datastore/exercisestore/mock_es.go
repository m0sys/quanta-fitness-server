package exercisestore

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/mhd53/quanta-fitness-server/internal/entity"
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
		ID:           string(s.lastID),
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

func (s *store) Update(eid string, updates entity.ExerciseUpdate) error {
	prev := s.exercises[parseInt64(eid)]
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

	s.exercises[parseInt64(eid)] = updated
	return nil
}

func (s *store) Delete(eid string) error {
	found := s.exercises[parseInt64(eid)]
	if isEmpty(found) {
		return errors.New("Not Found!")
	}

	delete(s.exercises, parseInt64(eid))
	return nil
}

func isEmpty(found entity.Exercise) bool {
	return found == (entity.Exercise{})
}

func (s *store) FindExerciseById(eid string) (entity.Exercise, bool, error) {
	found := s.exercises[parseInt64(eid)]

	if isEmpty(found) {
		return entity.Exercise{}, false, nil
	}

	return found, true, nil
}

func (s *store) FindAllExercisesByWID(wid string) ([]entity.Exercise, error) {
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

func parseInt64(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return val
}
