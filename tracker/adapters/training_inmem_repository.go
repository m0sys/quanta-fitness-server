package adapters

import (
	"errors"
	"time"

	"github.com/mhd53/quanta-fitness-server/manager/athlete"
	elg "github.com/mhd53/quanta-fitness-server/tracker/exerciselog"
	sl "github.com/mhd53/quanta-fitness-server/tracker/setlog"
	t "github.com/mhd53/quanta-fitness-server/tracker/training"
	wl "github.com/mhd53/quanta-fitness-server/tracker/workoutlog"
)

type repo struct {
	wlogs map[string]inRepoWorkoutLog
	elogs map[string]inRepoExerciseLog
	slogs map[string]inRepoSetLog
}

func NewInMemRepo() t.Repository {
	return &repo{
		wlogs: make(map[string]inRepoWorkoutLog),
		elogs: make(map[string]inRepoExerciseLog),
		slogs: make(map[string]inRepoSetLog),
	}
}
func (r *repo) StoreWorkoutLog(wlog wl.WorkoutLog) error {
	now := time.Now()
	data := inRepoWorkoutLog{
		ID:         wlog.ID(),
		AthleteID:  wlog.AthleteID(),
		Title:      wlog.Title(),
		Date:       wlog.Date(),
		CurrentPos: wlog.CurrentPos(),
		Completed:  wlog.Completed(),
		CreatedAt:  now,
		UpdatedAt:  now,
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
		Pos:          elog.Pos(),
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	r.elogs[elog.ID()] = data
	return nil
}

func (r *repo) StoreSetLog(slog sl.SetLog) error {
	metrics := slog.Metrics()
	now := time.Now()
	data := inRepoSetLog{
		ID:             slog.ID(),
		ExerciseLogID:  slog.ExerciseLogID(),
		ActualRepCount: metrics.ActualRepCount(),
		Dur:            float64(metrics.Dur()),
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	r.slogs[slog.ID()] = data
	return nil
}

func (r *repo) FindAllWorkoutLogsForAthlete(ath athlete.Athlete) ([]wl.WorkoutLog, error) {
	aid := ath.AthleteID()
	var wlogs []wl.WorkoutLog

	for _, val := range r.wlogs {
		if val.AthleteID == aid {
			wlog := wl.RestoreWorkoutLog(val.ID, val.AthleteID, val.Title, val.Date, val.CurrentPos, val.Completed)

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

			elog := elg.RestoreExerciseLog(val.ID, val.WorkoutLogID, val.Name, val.Completed, metrics, val.Pos)

			elogs = append(elogs, elog)
		}
	}

	return elogs, nil
}

func (r *repo) FindAllSetLogsForExerciseLog(elog elg.ExerciseLog) ([]sl.SetLog, error) {
	var slogs []sl.SetLog

	elid := elog.ID()

	for _, val := range r.slogs {
		if val.ExerciseLogID == elid {
			metrics := sl.NewMetrics(val.ActualRepCount, val.Dur)

			slog := sl.RestoreSetLog(val.ID, val.ExerciseLogID, metrics)

			slogs = append(slogs, slog)
		}
	}
	return slogs, nil
}

func (r *repo) FindWorkoutLogByID(wlog wl.WorkoutLog) (bool, error) {
	_, ok := r.wlogs[wlog.ID()]
	return ok, nil

}

func (r *repo) FindExerciseLogByID(elog elg.ExerciseLog) (bool, error) {
	_, ok := r.elogs[elog.ID()]
	return ok, nil

}

func (r *repo) RemoveWorkoutLog(wlog wl.WorkoutLog) error {
	delete(r.wlogs, wlog.ID())
	return nil
}

func (r *repo) RemoveExerciseLog(elog elg.ExerciseLog) error {
	delete(r.elogs, elog.ID())
	return nil
}

func (r *repo) RemoveSetLog(slog sl.SetLog) error {
	delete(r.slogs, slog.ID())
	return nil
}

func (r *repo) UpdateWorkoutLog(wlog wl.WorkoutLog) error {
	prev, ok := r.wlogs[wlog.ID()]
	if !ok {
		return errors.New("WorkoutLog not found!")
	}

	now := time.Now()
	data := inRepoWorkoutLog{
		ID:         wlog.ID(),
		AthleteID:  wlog.AthleteID(),
		Title:      wlog.Title(),
		CurrentPos: wlog.CurrentPos(),
		Completed:  wlog.Completed(),
		Date:       wlog.Date(),
		CreatedAt:  prev.CreatedAt,
		UpdatedAt:  now,
	}

	r.wlogs[wlog.ID()] = data
	return nil

}

type inRepoWorkoutLog struct {
	ID         string
	AthleteID  string
	Title      string
	CurrentPos int
	Completed  bool
	Date       time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
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
	Pos          int
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
