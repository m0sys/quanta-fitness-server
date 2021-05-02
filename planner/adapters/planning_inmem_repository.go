package adapters

import (
	"errors"
	"time"

	"github.com/mhd53/quanta-fitness-server/athlete"
	"github.com/mhd53/quanta-fitness-server/planner/exercise"
	p "github.com/mhd53/quanta-fitness-server/planner/planning"
	wp "github.com/mhd53/quanta-fitness-server/planner/workoutplan"
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

func (r *repo) StoreWorkoutPlan(wplan wp.WorkoutPlan, ath athlete.Athlete) error {
	now := time.Now()
	data := inRepoWorkoutPlan{
		ID:        wplan.ID(),
		AthleteID: ath.AthleteID(),
		Title:     wplan.Title(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	r.wplans[wplan.ID()] = data
	return nil
}

func (r *repo) FindWorkoutPlanByTitleAndAthleteID(title string, ath athlete.Athlete) (wp.WorkoutPlan, bool, error) {
	aid := ath.AthleteID()
	for _, val := range r.wplans {
		if val.AthleteID == aid && val.Title == title {
			found, err := wp.RestoreWorkoutPlan(val.ID, val.AthleteID, val.Title)
			if err != nil {
				return wp.WorkoutPlan{}, false, err
			}
			return found, true, nil
		}
	}

	return wp.WorkoutPlan{}, false, nil
}

func (r *repo) FindWorkoutPlanByID(wplan wp.WorkoutPlan) (bool, error) {
	_, ok := r.wplans[wplan.ID()]
	return ok, nil
}
func (r *repo) FindWorkoutPlanByIDAndAthleteID(wplan wp.WorkoutPlan, ath athlete.Athlete) (bool, error) {
	aid := ath.AthleteID()
	wpid := wplan.ID()

	for _, val := range r.wplans {
		if val.AthleteID == aid && val.ID == wpid {
			return true, nil
		}
	}

	return false, nil
}
func (r *repo) StoreExercise(wplan wp.WorkoutPlan, e exercise.Exercise, ath athlete.Athlete) error {
	now := time.Now()
	metrics := e.Metrics()
	data := inRepoExercise{
		ID:            e.ID(),
		WorkoutPlanID: wplan.ID(),
		AthleteID:     ath.AthleteID(),
		Name:          e.Name(),
		TargetRep:     metrics.TargetRep(),
		NumSets:       metrics.NumSets(),
		Weight:        float64(metrics.Weight()),
		RestDur:       float64(metrics.RestDur()),
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	r.exercises[e.ID()] = data
	return nil
}

func (r *repo) FindExerciseByID(e exercise.Exercise) (bool, error) {
	eid := e.ID()

	for _, val := range r.exercises {
		if val.ID == eid {
			return true, nil
		}
	}
	return false, nil
}

func (r *repo) FindExerciseByNameAndWorkoutPlanID(wplan wp.WorkoutPlan, e exercise.Exercise) (bool, error) {
	name := e.Name()
	wpid := wplan.ID()

	for _, val := range r.exercises {
		if val.Name == name && val.WorkoutPlanID == wpid {
			return true, nil
		}
	}

	return false, nil
}

func (r *repo) RemoveExercise(e exercise.Exercise) error {
	delete(r.exercises, e.ID())
	return nil
}

func (r *repo) UpdateWorkoutPlan(wplan wp.WorkoutPlan, title string) error {
	prev, ok := r.wplans[wplan.ID()]
	if !ok {
		return errors.New("WorkoutPlan not found!")
	}

	r.wplans[wplan.ID()] = inRepoWorkoutPlan{
		ID:        prev.ID,
		AthleteID: prev.AthleteID,
		Title:     title,
		CreatedAt: prev.CreatedAt,
		UpdatedAt: time.Now(),
	}
	return nil
}

func (r *repo) FindAllWorkoutPlansForAthlete(ath athlete.Athlete) ([]wp.WorkoutPlan, error) {
	aid := ath.AthleteID()
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

func (r *repo) FindAllExercisesForWorkoutPlan(wplan wp.WorkoutPlan) ([]exercise.Exercise, error) {
	wpid := wplan.ID()
	var exercises []exercise.Exercise
	for _, val := range r.exercises {
		if val.WorkoutPlanID == wpid {
			metrics, err := exercise.NewMetrics(val.TargetRep, val.NumSets, val.Weight, val.RestDur)
			if err != nil {
				return []exercise.Exercise{}, err
			}

			e, err := exercise.RestoreExercise(val.ID, val.WorkoutPlanID, val.AthleteID, val.Name, metrics)
			if err != nil {
				return []exercise.Exercise{}, err
			}

			exercises = append(exercises, e)
		}
	}

	return exercises, nil
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
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
