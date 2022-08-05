package adapters

import (
	"errors"
	"time"

	"github.com/m0sys/quanta-fitness-server/planner/exercise"
	e "github.com/m0sys/quanta-fitness-server/planner/exercise"
	p "github.com/m0sys/quanta-fitness-server/planner/planning"
	wp "github.com/m0sys/quanta-fitness-server/planner/workoutplan"
)

type repo struct {
	wplans    map[string]inRepoWorkoutPlan
	exercises map[string]inRepoExercise
}

func NewInMemRepo() p.Repository {
	return &repo{
		wplans:    make(map[string]inRepoWorkoutPlan),
		exercises: make(map[string]inRepoExercise),
	}
}

func (r *repo) StoreWorkoutPlan(wplan wp.WorkoutPlan) error {
	now := time.Now()
	data := inRepoWorkoutPlan{
		ID:        wplan.ID(),
		AthleteID: wplan.AthleteID(),
		Title:     wplan.Title(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	r.wplans[wplan.ID()] = data
	return nil
}

func (r *repo) FindWorkoutPlanByTitleAndAthleteID(wplan wp.WorkoutPlan) (bool, error) {
	aid := wplan.AthleteID()
	title := wplan.Title()

	for _, val := range r.wplans {
		if val.AthleteID == aid && val.Title == title {
			return true, nil
		}
	}

	return false, nil
}

func (r *repo) FindWorkoutPlanByID(id string) (wp.WorkoutPlan, bool, error) {
	val, ok := r.wplans[id]

	if !ok {
		return wp.WorkoutPlan{}, false, nil
	}

	wplan, err := wp.RestoreWorkoutPlan(val.ID, val.AthleteID, val.Title)
	if err != nil {
		return wp.WorkoutPlan{}, false, err
	}

	return wplan, ok, nil
}
func (r *repo) FindWorkoutPlanByIDAndAthleteID(wplan wp.WorkoutPlan) (bool, error) {
	aid := wplan.AthleteID()
	wpid := wplan.ID()

	for _, val := range r.wplans {
		if val.AthleteID == aid && val.ID == wpid {
			return true, nil
		}
	}

	return false, nil
}
func (r *repo) StoreExercise(wplan wp.WorkoutPlan, exercise e.Exercise) error {
	now := time.Now()
	metrics := exercise.Metrics()
	data := inRepoExercise{
		ID:            exercise.ID(),
		WorkoutPlanID: wplan.ID(),
		AthleteID:     wplan.AthleteID(),
		Name:          exercise.Name(),
		TargetRep:     metrics.TargetRep(),
		NumSets:       metrics.NumSets(),
		Weight:        float64(metrics.Weight()),
		RestDur:       float64(metrics.RestDur()),
		Pos:           exercise.Pos(),
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	r.exercises[exercise.ID()] = data
	return nil
}

func (r *repo) FindExerciseByID(id string) (e.Exercise, bool, error) {
	val, ok := r.exercises[id]
	if !ok {
		return e.Exercise{}, false, nil
	}

	exercise, err := r.restoreExercise(val)
	if err != nil {
		return e.Exercise{}, false, err
	}

	return exercise, ok, nil
}

func (r *repo) FindExerciseByNameAndWorkoutPlanID(wpid, name string) (bool, error) {
	for _, val := range r.exercises {
		if val.Name == name && val.WorkoutPlanID == wpid {
			return true, nil
		}
	}

	return false, nil
}

func (r *repo) RemoveExercise(exercise e.Exercise) error {
	delete(r.exercises, exercise.ID())
	return nil
}

func (r *repo) UpdateWorkoutPlan(wplan wp.WorkoutPlan) error {
	prev, ok := r.wplans[wplan.ID()]
	if !ok {
		return errors.New("WorkoutPlan not found!")
	}

	data := inRepoWorkoutPlan{
		ID:        prev.ID,
		AthleteID: prev.AthleteID,
		Title:     wplan.Title(),
		CreatedAt: prev.CreatedAt,
		UpdatedAt: time.Now(),
	}
	r.wplans[wplan.ID()] = data
	return nil
}

func (r *repo) FindAllWorkoutPlansForAthlete(aid string) ([]wp.WorkoutPlan, error) {
	var wplans []wp.WorkoutPlan
	for _, val := range r.wplans {
		if val.AthleteID == aid {
			wplan, err := wp.RestoreWorkoutPlan(val.ID, val.AthleteID, val.Title)
			if err != nil {
				return []wp.WorkoutPlan{}, err
			}

			wplans = append(wplans, wplan)
		}
	}

	return wplans, nil
}

func (r *repo) FindAllExercisesForWorkoutPlan(wplan wp.WorkoutPlan) ([]e.Exercise, error) {
	wpid := wplan.ID()
	var exercises []e.Exercise
	for _, val := range r.exercises {
		if val.WorkoutPlanID == wpid {
			exercise, err := r.restoreExercise(val)
			if err != nil {
				return []e.Exercise{}, err
			}

			exercises = append(exercises, exercise)
		}
	}

	return exercises, nil
}

func (r *repo) restoreExercise(val inRepoExercise) (e.Exercise, error) {
	metrics, err := exercise.NewMetrics(val.TargetRep, val.NumSets, val.Weight, val.RestDur)
	if err != nil {
		return e.Exercise{}, err
	}

	exercise, err := e.RestoreExercise(val.ID, val.WorkoutPlanID, val.AthleteID, val.Name, metrics, val.Pos)
	if err != nil {
		return e.Exercise{}, err
	}

	return exercise, nil
}

func (r *repo) UpdateExercise(exercise e.Exercise) error {
	prev, ok := r.exercises[exercise.ID()]
	if !ok {
		return errors.New("Exercise not found!")
	}

	now := time.Now()
	metrics := exercise.Metrics()
	data := inRepoExercise{
		ID:            exercise.ID(),
		WorkoutPlanID: exercise.WorkoutPlanID(),
		AthleteID:     exercise.AthleteID(),
		Name:          exercise.Name(),
		TargetRep:     metrics.TargetRep(),
		NumSets:       metrics.NumSets(),
		Weight:        float64(metrics.Weight()),
		RestDur:       float64(metrics.RestDur()),
		CreatedAt:     prev.CreatedAt,
		UpdatedAt:     now,
	}

	r.exercises[exercise.ID()] = data
	return nil
}

func (r *repo) RemoveWorkoutPlan(wplan wp.WorkoutPlan) error {
	delete(r.wplans, wplan.ID())
	return nil
}

type inRepoWorkoutPlan struct {
	ID        string
	AthleteID string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type inRepoExercise struct {
	ID            string
	WorkoutPlanID string
	AthleteID     string
	Name          string
	TargetRep     int
	NumSets       int
	Weight        float64
	RestDur       float64
	Pos           int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
