package workoutstore

import (
	"errors"
	"log"
	"strconv"
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
		ID:          string(s.lastID),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.workouts[s.lastID] = created
	s.lastID += 1

	return created, nil
}

func (s *store) Update(wid string, updates entity.BaseWorkout) error {
	prev := s.workouts[parseInt64(wid)]
	updated := entity.Workout{
		BaseWorkout: updates,
		ID:          prev.ID,
		CreatedAt:   prev.CreatedAt,
		UpdatedAt:   time.Now(),
	}

	s.workouts[parseInt64(wid)] = updated

	return nil
}

func (s *store) FindWorkoutById(wid string) (entity.Workout, bool, error) {
	found := s.workouts[parseInt64(wid)]

	if isEmpty(found) {
		return entity.Workout{}, false, nil
	}

	return found, true, nil
}

func isEmpty(found entity.Workout) bool {
	return found == (entity.Workout{})
}

func (s *store) DeleteWorkout(wid string) error {
	found := s.workouts[parseInt64(wid)]
	if isEmpty(found) {
		return errors.New("Not Found!")
	}

	delete(s.workouts, parseInt64(wid))
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

func parseInt64(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return val
}
