package adapters

import (
	"context"

	"database/sql"

	"github.com/google/uuid"

	db "github.com/mhd53/quanta-fitness-server/internal/db/sqlc"
	"github.com/mhd53/quanta-fitness-server/manager/athlete"
	e "github.com/mhd53/quanta-fitness-server/planner/exercise"
	wp "github.com/mhd53/quanta-fitness-server/planner/workoutplan"
)

type PsqlPlanningRepository struct {
	store *db.Store
}

func NewPSQLRepo(store *db.Store) PsqlPlanningRepository {
	return PsqlPlanningRepository{
		store: store,
	}
}

func (r *PsqlPlanningRepository) StoreWorkoutPlan(wplan wp.WorkoutPlan, ath athlete.Athlete) error {
	id, err := uuid.Parse(wplan.ID())
	if err != nil {
		return err
	}

	aid, err := uuid.Parse(wplan.AthleteID())
	if err != nil {
		return err
	}

	arg := db.StoreWorkoutPlanParams{
		ID:    id,
		Aid:   aid,
		Title: wplan.Title(),
	}

	ctx := context.Background()
	return r.store.StoreWorkoutPlan(ctx, arg)
}

func (r *PsqlPlanningRepository) FindWorkoutPlanByTitleAndAthleteID(wplan wp.WorkoutPlan, ath athlete.Athlete) (bool, error) {
	aid, err := uuid.Parse(wplan.AthleteID())
	if err != nil {
		return false, err
	}

	arg := db.FindWorkoutPlanByTitleAndAthleteIDParams{
		Title: wplan.Title(),
		Aid:   aid,
	}

	ctx := context.Background()
	_, err = r.store.FindWorkoutPlanByTitleAndAthleteID(ctx, arg)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *PsqlPlanningRepository) FindWorkoutPlanByID(wplan wp.WorkoutPlan) (bool, error) {
	id, err := uuid.Parse(wplan.ID())
	if err != nil {
		return false, err
	}

	ctx := context.Background()
	_, err = r.store.FindWorkoutPlanByID(ctx, id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *PsqlPlanningRepository) FindWorkoutPlanByIDAndAthleteID(wplan wp.WorkoutPlan, ath athlete.Athlete) (bool, error) {
	id, err := uuid.Parse(wplan.ID())
	if err != nil {
		return false, err
	}

	aid, err := uuid.Parse(wplan.AthleteID())
	if err != nil {
		return false, err
	}

	arg := db.FindWorkoutPlanByIDAndAthleteIDParams{
		ID:  id,
		Aid: aid,
	}

	ctx := context.Background()
	_, err = r.store.FindWorkoutPlanByIDAndAthleteID(ctx, arg)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *PsqlPlanningRepository) StoreExercise(
	wplan wp.WorkoutPlan,
	exercise e.Exercise,
	ath athlete.Athlete,
) error {
	id, err := uuid.Parse(exercise.ID())
	if err != nil {
		return err
	}

	wpid, err := uuid.Parse(wplan.ID())
	if err != nil {
		return err
	}

	aid, err := uuid.Parse(wplan.AthleteID())
	if err != nil {
		return err
	}

	metrics := exercise.Metrics()

	arg := db.StoreExerciseParams{
		ID:           id,
		Aid:          aid,
		Wpid:         wpid,
		Name:         exercise.Name(),
		TargetRep:    int32(metrics.TargetRep()),
		NumSets:      int32(metrics.NumSets()),
		Weight:       float64(metrics.Weight()),
		RestDuration: float64(metrics.RestDur()),
		Pos:          int32(exercise.Pos()),
	}

	ctx := context.Background()
	return r.store.StoreExercise(ctx, arg)
}

func (r *PsqlPlanningRepository) FindExerciseByID(exercise e.Exercise) (bool, error) {
	id, err := uuid.Parse(exercise.ID())
	if err != nil {
		return false, err
	}

	ctx := context.Background()
	_, err = r.store.FindExerciseByID(ctx, id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *PsqlPlanningRepository) FindExerciseByNameAndWorkoutPlanID(wplan wp.WorkoutPlan, exercise e.Exercise) (bool, error) {
	wpid, err := uuid.Parse(wplan.ID())
	if err != nil {
		return false, err
	}

	arg := db.FindExerciseByNameAndWorkoutPlanIDParams{
		Name: exercise.Name(),
		Wpid: wpid,
	}

	ctx := context.Background()
	_, err = r.store.FindExerciseByNameAndWorkoutPlanID(ctx, arg)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *PsqlPlanningRepository) RemoveExercise(exercise e.Exercise) error {
	id, err := uuid.Parse(exercise.ID())
	if err != nil {
		return err
	}

	ctx := context.Background()
	return r.store.RemoveExercise(ctx, id)
}

func (r *PsqlPlanningRepository) UpdateWorkoutPlan(wplan wp.WorkoutPlan) error {
	id, err := uuid.Parse(wplan.ID())
	if err != nil {
		return err
	}

	arg := db.UpdateWorkoutPlanParams{
		ID:    id,
		Title: wplan.Title(),
	}

	ctx := context.Background()
	return r.store.UpdateWorkoutPlan(ctx, arg)
}

func (r *PsqlPlanningRepository) FindAllWorkoutPlansForAthlete(ath athlete.Athlete) ([]wp.WorkoutPlan, error) {
	var wplans []wp.WorkoutPlan
	id, err := uuid.Parse(ath.AthleteID())
	if err != nil {
		return wplans, err
	}

	ctx := context.Background()
	dbWplans, err := r.store.FindAllWorkoutPlansForAthlete(ctx, id)

	for _, dbplan := range dbWplans {
		wplan, err := wp.RestoreWorkoutPlan(dbplan.ID.String(), dbplan.Aid.String(), dbplan.Title)
		if err != nil {
			return []wp.WorkoutPlan{}, err
		}
		wplans = append(wplans, wplan)
	}

	return wplans, nil
}

func (r *PsqlPlanningRepository) FindAllExercisesForWorkoutPlan(wplan wp.WorkoutPlan) ([]e.Exercise, error) {
	var exercises []e.Exercise
	id, err := uuid.Parse(wplan.ID())
	if err != nil {
		return exercises, err
	}

	ctx := context.Background()
	dbexs, err := r.store.FindAllExercisesForWorkoutPlan(ctx, id)

	for _, dbex := range dbexs {
		metrics, err := e.NewMetrics(int(dbex.TargetRep), int(dbex.NumSets), dbex.Weight, dbex.RestDuration)
		if err != nil {
			return []e.Exercise{}, err
		}

		exercise, err := e.RestoreExercise(dbex.ID.String(), dbex.Wpid.String(), dbex.Aid.String(), dbex.Name, metrics, int(dbex.Pos))
		if err != nil {
			return []e.Exercise{}, err
		}

		exercises = append(exercises, exercise)
	}
	return exercises, nil
}

func (r *PsqlPlanningRepository) UpdateExercise(exercise e.Exercise) error {
	id, err := uuid.Parse(exercise.ID())
	if err != nil {
		return err
	}

	metrics := exercise.Metrics()
	ctx := context.Background()
	arg := db.UpdateExerciseParams{
		ID:           id,
		Name:         exercise.Name(),
		TargetRep:    int32(metrics.TargetRep()),
		NumSets:      int32(metrics.NumSets()),
		Weight:       float64(metrics.Weight()),
		RestDuration: float64(metrics.RestDur()),
		Pos:          int32(exercise.Pos()),
	}
	return r.store.UpdateExercise(ctx, arg)
}

func (r *PsqlPlanningRepository) RemoveWorkoutPlan(wplan wp.WorkoutPlan) error {
	id, err := uuid.Parse(wplan.ID())
	if err != nil {
		return err
	}

	ctx := context.Background()
	return r.store.RemoveWorkoutPlan(ctx, id)
}
