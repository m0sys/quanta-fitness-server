package adapters

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	db "github.com/mhd53/quanta-fitness-server/internal/db/sqlc"
	"github.com/mhd53/quanta-fitness-server/manager/athlete"
	el "github.com/mhd53/quanta-fitness-server/tracker/exerciselog"
	sl "github.com/mhd53/quanta-fitness-server/tracker/setlog"
	wl "github.com/mhd53/quanta-fitness-server/tracker/workoutlog"
)

type PsqlTrainingRepository struct {
	store *db.Store
}

func NewPSQLRepo(store *db.Store) PsqlTrainingRepository {
	return PsqlTrainingRepository{
		store: store,
	}
}

func (r *PsqlTrainingRepository) StoreWorkoutLog(wlog wl.WorkoutLog) error {
	id, err := uuid.Parse(wlog.ID())
	if err != nil {
		return err
	}

	aid, err := uuid.Parse(wlog.AthleteID())
	if err != nil {
		return err
	}

	arg := db.StoreWorkoutLogParams{
		ID:         id,
		Aid:        aid,
		Title:      wlog.Title(),
		Date:       wlog.Date(),
		CurrentPos: int32(wlog.CurrentPos()),
		Completed:  wlog.Completed(),
	}

	return r.store.StoreWorkoutLog(context.Background(), arg)
}

func (r *PsqlTrainingRepository) StoreExerciseLog(elog el.ExerciseLog) error {
	id, err := uuid.Parse(elog.ID())
	if err != nil {
		return err
	}

	wlid, err := uuid.Parse(elog.WorkoutLogID())
	if err != nil {
		return err
	}

	metrics := elog.Metrics()
	arg := db.StoreExerciseLogParams{
		ID:           id,
		Wlid:         wlid,
		Name:         elog.Name(),
		TargetRep:    int32(metrics.TargetRep()),
		NumSets:      int32(metrics.NumSets()),
		Weight:       float64(metrics.Weight()),
		RestDuration: float64(metrics.RestDur()),
		Completed:    elog.Completed(),
		Pos:          int32(elog.Pos()),
	}

	return r.store.StoreExerciseLog(context.Background(), arg)
}

func (r *PsqlTrainingRepository) StoreSetLog(slog sl.SetLog) error {
	id, err := uuid.Parse(slog.ID())
	if err != nil {
		return err
	}

	elid, err := uuid.Parse(slog.ExerciseLogID())
	if err != nil {
		return err
	}

	metrics := slog.Metrics()
	arg := db.StoreSetLogParams{
		ID:             id,
		Elid:           elid,
		ActualRepCount: int32(metrics.ActualRepCount()),
		Duration:       float64(metrics.Dur()),
	}

	return r.store.StoreSetLog(context.Background(), arg)
}

func (r *PsqlTrainingRepository) FindAllWorkoutLogsForAthlete(ath athlete.Athlete) ([]wl.WorkoutLog, error) {
	var wlogs []wl.WorkoutLog
	aid, err := uuid.Parse(ath.AthleteID())
	if err != nil {
		return wlogs, err
	}

	dbWlogs, err := r.store.FindAllWorkoutLogsForAthlete(context.Background(), aid)
	if err != nil {
		return wlogs, err
	}

	for _, dblog := range dbWlogs {
		wlog := wl.RestoreWorkoutLog(dblog.ID.String(), dblog.Aid.String(), dblog.Title, dblog.Date, int(dblog.CurrentPos), dblog.Completed)
		if err != nil {
			return []wl.WorkoutLog{}, err
		}
		wlogs = append(wlogs, wlog)
	}

	return wlogs, nil
}

func (r *PsqlTrainingRepository) FindAllExerciseLogsForWorkoutLog(wlog wl.WorkoutLog) ([]el.ExerciseLog, error) {
	var elogs []el.ExerciseLog
	wlid, err := uuid.Parse(wlog.ID())
	if err != nil {
		return elogs, err
	}

	dbElogs, err := r.store.FindAllExerciseLogsForWorkoutLog(context.Background(), wlid)
	if err != nil {
		return elogs, err
	}

	for _, dblog := range dbElogs {
		metrics := el.NewMetrics(int(dblog.TargetRep), int(dblog.NumSets), dblog.Weight, dblog.RestDuration)
		elog := el.RestoreExerciseLog(dblog.ID.String(), dblog.Wlid.String(), dblog.Name, dblog.Completed, metrics, int(dblog.Pos))
		if err != nil {
			return []el.ExerciseLog{}, err
		}
		elogs = append(elogs, elog)
	}
	return elogs, nil
}

func (r *PsqlTrainingRepository) FindAllSetLogsForExerciseLog(elog el.ExerciseLog) ([]sl.SetLog, error) {
	var slogs []sl.SetLog
	elid, err := uuid.Parse(elog.ID())
	if err != nil {
		return slogs, err
	}

	dbSlogs, err := r.store.FindAllSetLogsForExerciseLog(context.Background(), elid)
	if err != nil {
		return slogs, err
	}

	for _, dblog := range dbSlogs {
		metrics := sl.NewMetrics(int(dblog.ActualRepCount), dblog.Duration)
		slog := sl.RestoreSetLog(dblog.ID.String(), dblog.Elid.String(), metrics)
		if err != nil {
			return []sl.SetLog{}, err
		}
		slogs = append(slogs, slog)
	}
	return slogs, nil
}

func (r *PsqlTrainingRepository) FindWorkoutLogByID(wlog wl.WorkoutLog) (bool, error) {
	id, err := uuid.Parse(wlog.ID())
	if err != nil {
		return false, err
	}

	ctx := context.Background()
	_, err = r.store.FindWorkoutLogByID(ctx, id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *PsqlTrainingRepository) FindExerciseLogByID(elog el.ExerciseLog) (bool, error) {
	id, err := uuid.Parse(elog.ID())
	if err != nil {
		return false, err
	}

	ctx := context.Background()
	_, err = r.store.FindExerciseLogByID(ctx, id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *PsqlTrainingRepository) UpdateWorkoutLog(wlog wl.WorkoutLog) error {
	id, err := uuid.Parse(wlog.ID())
	if err != nil {
		return err
	}
	arg := db.UpdateWorkoutLogParams{
		ID:         id,
		CurrentPos: int32(wlog.CurrentPos()),
		Completed:  wlog.Completed(),
	}

	ctx := context.Background()
	return r.store.UpdateWorkoutLog(ctx, arg)
}

func (r *PsqlTrainingRepository) RemoveWorkoutLog(wlog wl.WorkoutLog) error {
	id, err := uuid.Parse(wlog.ID())
	if err != nil {
		return err
	}

	ctx := context.Background()
	return r.store.RemoveWorkoutLog(ctx, id)
}

func (r *PsqlTrainingRepository) RemoveExerciseLog(elog el.ExerciseLog) error {
	id, err := uuid.Parse(elog.ID())
	if err != nil {
		return err
	}

	ctx := context.Background()
	return r.store.RemoveExerciseLog(ctx, id)
}

func (r *PsqlTrainingRepository) RemoveSetLog(slog sl.SetLog) error {
	id, err := uuid.Parse(slog.ID())
	if err != nil {
		return err
	}

	ctx := context.Background()
	return r.store.RemoveSetLog(ctx, id)
}
