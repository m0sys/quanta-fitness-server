package trackerrepo

import (
	"time"

	"github.com/mhd53/quanta-fitness-server/athlete"
	tckr "github.com/mhd53/quanta-fitness-server/tracker"
	wl "github.com/mhd53/quanta-fitness-server/workoutlog"
)

type repo struct {
	wlogs     map[string]inRepoWorkoutLog
	exercises map[string]inRepoExercise
}

// NewTrackerRepo create new in memory Tracker Repository.
func NewTrackerRepo() tckr.Repository {
	return &repo{
		wlogs:     make(map[string]inRepoWorkoutLog),
		exercises: make(map[string]inRepoExercise),
	}
}

func (r *repo) StoreWorkoutLog(wlog wl.WorkoutLog, ath athlete.Athlete) (tckr.WorkoutLogRes, error) {
	createTime := time.Now()
	rWlog := inRepoWorkoutLog{
		LogID:     wlog.LogID,
		AthleteID: ath.AthleteID,
		Title:     wlog.Title,
		Date:      wlog.Date,
		CreatedAt: createTime,
		UpdatedAt: createTime,
	}
	r.wlogs[wlog.LogID] = rWlog

	return tckr.WorkoutLogRes{
		LogID:     rWlog.LogID,
		Title:     rWlog.Title,
		Date:      rWlog.Date,
		CreatedAt: rWlog.CreatedAt,
		UpdatedAt: rWlog.UpdatedAt,
	}, nil
}

// FindWorkoutLogByID find WorkoutLog in memory Tracker Repository by LogID.
func (r *repo) FindWorkoutLogByID(id string) (tckr.WorkoutLogRes, bool, error) {
	var rWlog inRepoWorkoutLog
	var found bool

	for k, val := range r.wlogs {
		if k == id {
			rWlog = val
			found = true
		}
	}

	if !found {
		return tckr.WorkoutLogRes{}, false, nil
	}

	return tckr.WorkoutLogRes{
		LogID:     rWlog.LogID,
		Title:     rWlog.Title,
		Date:      rWlog.Date,
		CreatedAt: rWlog.CreatedAt,
		UpdatedAt: rWlog.UpdatedAt,
	}, true, nil
}

func (r *repo) FindAllExercisesForWorkoutLog(wlog wl.WorkoutLog) ([]tckr.ExerciseRes, error) {
	var exercises []tckr.ExerciseRes

	for _, val := range r.exercises {
		if val.LogID == wlog.LogID {
			found := tckr.ExerciseRes{
				ExerciseID: val.ExerciseID,
				Name:       val.Name,
				Weight:     val.Weight,
				TargetRep:  val.TargetRep,
				CreatedAt:  val.CreatedAt,
				UpdatedAt:  val.UpdatedAt,
			}
			exercises = append(exercises, found)
		}
	}

	return exercises, nil

}

func (r *repo) AddExerciseToWorkoutLog(wlog wl.WorkoutLog, exercise wl.Exercise) (tckr.ExerciseRes, error) {
	createTime := time.Now()
	rExercise := inRepoExercise{
		ExerciseID: exercise.ExerciseID,
		LogID:      wlog.LogID,
		Name:       exercise.Name,
		Weight:     exercise.Weight,
		TargetRep:  exercise.TargetRep,
		RestTime:   exercise.RestTime,
		CreatedAt:  createTime,
		UpdatedAt:  createTime,
	}
	r.exercises[exercise.ExerciseID] = rExercise

	return tckr.ExerciseRes{
		ExerciseID: rExercise.ExerciseID,
		Name:       rExercise.Name,
		Weight:     rExercise.Weight,
		TargetRep:  rExercise.TargetRep,
		RestTime:   rExercise.RestTime,
		CreatedAt:  rExercise.CreatedAt,
		UpdatedAt:  rExercise.UpdatedAt,
	}, nil
}

type inRepoWorkoutLog struct {
	LogID     string
	AthleteID string
	Title     string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type inRepoExercise struct {
	ExerciseID string
	LogID      string
	Name       string
	Weight     float64
	TargetRep  int
	RestTime   float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
