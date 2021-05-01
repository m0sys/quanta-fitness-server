package planning

import (
	"github.com/mhd53/quanta-fitness-server/athlete"
	"github.com/mhd53/quanta-fitness-server/planner/exercise"
	wp "github.com/mhd53/quanta-fitness-server/planner/workoutplan"
)

type Repository interface {
	StoreWorkoutPlan(wplan wp.WorkoutPlan, ath athlete.Athlete) error
	FindWorkoutPlanByTitleAndAthleteID(title string, ath athlete.Athlete) (wp.WorkoutPlan, bool, error)
	FindWorkoutPlanByID(wplan wp.WorkoutPlan) (bool, error)
	FindWorkoutPlanByIDAndAthleteID(wplan wp.WorkoutPlan, ath athlete.Athlete) (bool, error)
	StoreExercise(wplan wp.WorkoutPlan, e exercise.Exercise, ath athlete.Athlete) error
}
