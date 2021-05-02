package planning

import (
	"github.com/mhd53/quanta-fitness-server/athlete"
	e "github.com/mhd53/quanta-fitness-server/planner/exercise"
	wp "github.com/mhd53/quanta-fitness-server/planner/workoutplan"
)

type Repository interface {
	StoreWorkoutPlan(wplan wp.WorkoutPlan, ath athlete.Athlete) error
	FindWorkoutPlanByTitleAndAthleteID(wplan wp.WorkoutPlan, ath athlete.Athlete) (bool, error)
	FindWorkoutPlanByID(wplan wp.WorkoutPlan) (bool, error)
	FindWorkoutPlanByIDAndAthleteID(wplan wp.WorkoutPlan, ath athlete.Athlete) (bool, error)
	StoreExercise(wplan wp.WorkoutPlan, exercise e.Exercise, ath athlete.Athlete) error
	FindExerciseByID(exercise e.Exercise) (bool, error)
	FindExerciseByNameAndWorkoutPlanID(wplan wp.WorkoutPlan, exercise e.Exercise) (bool, error)
	RemoveExercise(exercise e.Exercise) error
	UpdateWorkoutPlan(wplan wp.WorkoutPlan) error
	FindAllWorkoutPlansForAthlete(ath athlete.Athlete) ([]wp.WorkoutPlan, error)
	FindAllExercisesForWorkoutPlan(wplan wp.WorkoutPlan) ([]e.Exercise, error)
	UpdateExercise(exercise e.Exercise) error
	RemoveWorkoutPlan(wplan wp.WorkoutPlan) error
}
