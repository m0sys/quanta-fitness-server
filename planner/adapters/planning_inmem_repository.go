package adapters

import (
	"time"

	"github.com/mhd53/quanta-fitness-server/athlete"
	p "github.com/mhd53/quanta-fitness-server/planner/planning"
	wp "github.com/mhd53/quanta-fitness-server/planner/workoutplan"
)

type repo struct {
	wplans map[string]inRepoWorkoutPlan
	// exercises map[string]inRepoExercise
}

func NewInMemRepo() p.Repository {
	return &repo{
		wplans: make(map[string]inRepoWorkoutPlan),
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
		if val.AthleteID == aid {
			if val.Title == title {
				found, err := wp.RestoreWorkoutPlan(val.ID, val.Title)
				if err != nil {
					return wp.WorkoutPlan{}, false, err
				}
				return found, true, nil
			}
		}
	}

	return wp.WorkoutPlan{}, false, nil
}

type inRepoWorkoutPlan struct {
	ID        string
	AthleteID string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
