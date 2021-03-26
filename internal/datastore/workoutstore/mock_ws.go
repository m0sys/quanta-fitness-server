package workoutstore

import (
	"errors"
	"time"

	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type store struct {
	workouts map[int64]entity.Workout
	lastID   int64
}

func NewMockWorkoutStore() WorkoutStore {
	return &store{workouts: make(map[int64]entity.Workout)}
}

func (s *store) Save(workout entity.BaseWorkout) (entity.Workout, error) {
	created := entity.Workout{
		BaseWorkout: workout,
		ID:          s.lastID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.workouts[s.lastID] = created
	s.lastID += 1

	return created, nil
}

func (s *store) Update(wid int64, updates entity.BaseWorkout) error {
	prev := s.workouts[wid]
	updated := entity.Workout{
		BaseWorkout: updates,
		ID:          prev.ID,
		CreatedAt:   prev.CreatedAt,
		UpdatedAt:   time.Now(),
	}

	s.workouts[wid] = updated

	return nil
}

func (s *store) FindWorkoutById(wid int64) (entity.Workout, bool, error) {
	found := s.workouts[wid]

	if isEmpty(found) {
		return entity.Workout{}, false, nil
	}

	return found, true, nil
}

func isEmpty(found entity.Workout) bool {
	return found == (entity.Workout{})
}

func (s *store) DeleteWorkout(wid int64) error {
	found := s.workouts[wid]
	if isEmpty(found) {
		return errors.New("Not Found!")
	}

	delete(s.workouts, wid)
	return nil

}

func (s *store) FindAllWorkoutsByUname(uname string) ([]entity.Workout, error) {
	var entries []entity.Workout

	for k := range s.workouts {
		entry := s.workouts[k]
		if entry.Username == uname {
			entries = append(entries, entry)
		}
	}

	return entries, nil
}
