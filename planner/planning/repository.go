package planning

import (
	"github.com/mhd53/quanta-fitness-server/athlete"
	wp "github.com/mhd53/quanta-fitness-server/planner/workoutplan"
)

type Repository interface {
	StoreWorkoutPlan(wplan wp.WorkoutPlan, ath athlete.Athlete) error
	FindWorkoutPlanByTitleAndAthleteID(title string, ath athlete.Athlete) (wp.WorkoutPlan, bool, error)
}
