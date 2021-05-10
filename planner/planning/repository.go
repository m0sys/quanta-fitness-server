package planning

import (
	e "github.com/mhd53/quanta-fitness-server/planner/exercise"
	wp "github.com/mhd53/quanta-fitness-server/planner/workoutplan"
)

type Repository interface {
	StoreWorkoutPlan(wplan wp.WorkoutPlan) error
	FindWorkoutPlanByTitleAndAthleteID(wplan wp.WorkoutPlan) (bool, error)
	FindWorkoutPlanByID(id string) (wp.WorkoutPlan, bool, error)
	FindWorkoutPlanByIDAndAthleteID(wplan wp.WorkoutPlan) (bool, error)
	StoreExercise(wplan wp.WorkoutPlan, exercise e.Exercise) error
	FindExerciseByID(id string) (e.Exercise, bool, error)
	FindExerciseByNameAndWorkoutPlanID(wpid, name string) (bool, error)
	RemoveExercise(exercise e.Exercise) error
	UpdateWorkoutPlan(wplan wp.WorkoutPlan) error
	FindAllWorkoutPlansForAthlete(aid string) ([]wp.WorkoutPlan, error)
	FindAllExercisesForWorkoutPlan(wplan wp.WorkoutPlan) ([]e.Exercise, error)
	UpdateExercise(exercise e.Exercise) error
	RemoveWorkoutPlan(wplan wp.WorkoutPlan) error
}
