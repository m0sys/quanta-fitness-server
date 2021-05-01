package planning

import (
	"fmt"
	"log"

	"github.com/mhd53/quanta-fitness-server/athlete"
	"github.com/mhd53/quanta-fitness-server/planner/exercise"
	wp "github.com/mhd53/quanta-fitness-server/planner/workoutplan"
)

const (
	errSlur = "planning"
)

var (
	errInternal                     = fmt.Errorf("%s: internal error", errSlur)
	ErrIdentialTitle                = fmt.Errorf("%s: WorkoutPlan with identical title already exists", errSlur)
	ErrWorkoutPlanAlreadyExists     = fmt.Errorf("%s: WorkoutPlan already exists", errSlur)
	ErrUnauthorizedAccess           = fmt.Errorf("%s: unauthorized access", errSlur)
	ErrWorkoutPlanNotFound          = fmt.Errorf("%s: WorkoutPlan not found", errSlur)
	ErrIdentialName                 = fmt.Errorf("%s: Exercise with identical name already in WorkoutPlan", errSlur)
	ErrExerciseAlreadyInWorkoutPlan = fmt.Errorf("%s: Exercise already in WorkoutPlan", errSlur)
)

type PlanningService struct {
	repo Repository
}

func NewPlanningService(repository Repository) PlanningService {
	return PlanningService{repo: repository}
}

func (p PlanningService) CreateNewWorkoutPlan(ath athlete.Athlete, wplan wp.WorkoutPlan) error {
	found, err := p.repo.FindWorkoutPlanByID(wplan)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	if found {
		return ErrWorkoutPlanAlreadyExists
	}

	_, found, err = p.repo.FindWorkoutPlanByTitleAndAthleteID(wplan.Title(), ath)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	if found {
		return ErrIdentialTitle
	}

	err = p.repo.StoreWorkoutPlan(wplan, ath)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	return nil
}

func (p PlanningService) AddNewExerciseToWorkoutPlan(
	ath athlete.Athlete,
	wplan wp.WorkoutPlan,
	exercise exercise.Exercise,
) error {
	found, err := p.repo.FindWorkoutPlanByID(wplan)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	if !found {
		return ErrWorkoutPlanNotFound
	}

	found, err = p.repo.FindWorkoutPlanByIDAndAthleteID(wplan, ath)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	if !found {
		return ErrUnauthorizedAccess
	}

	found, err = p.repo.FindExerciseByID(exercise)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	if found {
		return ErrExerciseAlreadyInWorkoutPlan
	}

	found, err = p.repo.FindExerciseByNameAndWorkoutPlanID(wplan, exercise)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	if found {
		return ErrIdentialName
	}

	err = p.repo.StoreExercise(wplan, exercise, ath)
	return nil
}
