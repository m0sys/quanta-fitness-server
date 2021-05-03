package adapters

import (
	"time"

	"github.com/mhd53/quanta-fitness-server/manager/athlete"
	elg "github.com/mhd53/quanta-fitness-server/tracker/exerciselog"
	sl "github.com/mhd53/quanta-fitness-server/tracker/setlog"
	t "github.com/mhd53/quanta-fitness-server/tracker/training"
	wl "github.com/mhd53/quanta-fitness-server/tracker/workoutlog"
)

type repo struct {
	wlogs   map[string]inRepoWorkoutLog
	elogs   map[string]inRepoExerciseLog
	setlogs map[string]inRepoSetLog
}

func NewInMemRepo() t.Repository {
	return &repo{
		wlogs:   make(map[string]inRepoWorkoutLog),
		elogs:   make(map[string]inRepoExerciseLog),
		setlogs: make(map[string]inRepoSetLog),
	}
}
func (r *repo) StoreWorkoutLog(wlog wl.WorkoutLog) error {
	now := time.Now()
	data := inRepoWorkoutLog{
		ID:        wlog.ID(),
		AthleteID: wlog.AthleteID(),
		Title:     wlog.Title(),
		Date:      wlog.Date(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	r.wlogs[wlog.ID()] = data
	return nil
}

func (r *repo) StoreExerciseLog(elog elg.ExerciseLog) error {
	metrics := elog.Metrics()
	now := time.Now()
	data := inRepoExerciseLog{
		ID:           elog.ID(),
		WorkoutLogID: elog.WorkoutLogID(),
		Name:         elog.Name(),
		TargetRep:    metrics.TargetRep(),
		NumSets:      metrics.NumSets(),
		Weight:       float64(metrics.Weight()),
		RestDur:      float64(metrics.RestDur()),
		Completed:    elog.Completed(),
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	r.elogs[elog.ID()] = data
	return nil
}

func (r *repo) StoreSetLog(setlog sl.SetLog) error {
	return nil
}

func (r *repo) RemoveWorkoutLog(wlog wl.WorkoutLog) error {
	return nil
}

func (r *repo) FindAllWorkoutLogsForAthlete(ath athlete.Athlete) ([]wl.WorkoutLog, error) {
	aid := ath.AthleteID()
	var wlogs []wl.WorkoutLog

	for _, val := range r.wlogs {
		if val.AthleteID == aid {
			wlog := wl.RestoreWorkoutLog(val.ID, val.AthleteID, val.Title, val.Date, val.Completed)

			wlogs = append(wlogs, wlog)
		}
	}

	return wlogs, nil
}
func (r *repo) FindAllExerciseLogsForWorkoutLog(wlog wl.WorkoutLog) ([]elg.ExerciseLog, error) {
	var elogs []elg.ExerciseLog
	wlid := wlog.ID()

	for _, val := range r.elogs {
		if val.WorkoutLogID == wlid {
			metrics := elg.NewMetrics(val.TargetRep, val.NumSets, val.Weight, val.RestDur)

			elog := elg.RestoreExerciseLog(val.ID, val.WorkoutLogID, val.Name, val.Completed, metrics)

			elogs = append(elogs, elog)
		}
	}

	return elogs, nil
}

func (r *repo) FindAllSetLogsForExerciseLog(elog elg.ExerciseLog) ([]sl.SetLog, error) {
	var slogs []sl.SetLog
	return slogs, nil
}

func (r *repo) FindWorkoutLogByID(wlog wl.WorkoutLog) (bool, error) {
	_, ok := r.wlogs[wlog.ID()]
	return ok, nil

}

type inRepoWorkoutLog struct {
	ID        string
	AthleteID string
	Title     string
	Completed bool
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type inRepoExerciseLog struct {
	ID           string
	WorkoutLogID string
	Name         string
	TargetRep    int
	NumSets      int
	Weight       float64
	RestDur      float64
	Completed    bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type inRepoSetLog struct {
	ID             string
	ExerciseLogID  string
	ActualRepCount int
	Dur            float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}