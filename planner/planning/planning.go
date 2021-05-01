package planning

import (
	"log"

	"github.com/mhd53/quanta-fitness-server/athlete"
	e "github.com/mhd53/quanta-fitness-server/planner/exercise"
	wp "github.com/mhd53/quanta-fitness-server/planner/workoutplan"
)

type PlanningService struct {
	repo Repository
}

func NewPlanningService(repository Repository) PlanningService {
	return PlanningService{repo: repository}
}

func (p PlanningService) CreateNewWorkoutPlan(ath athlete.Athlete, title string) (wp.WorkoutPlan, error) {
	wplan, err := wp.NewWorkoutPlan(ath.AthleteID(), title)
	if err != nil {
		return wp.WorkoutPlan{}, err
	}

	_, found, err := p.repo.FindWorkoutPlanByTitleAndAthleteID(wplan.Title(), ath)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return wp.WorkoutPlan{}, errInternal
	}

	if found {
		return wp.WorkoutPlan{}, ErrIdentialTitle
	}

	err = p.repo.StoreWorkoutPlan(wplan, ath)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return wp.WorkoutPlan{}, errInternal
	}

	return wplan, nil
}

func (p PlanningService) AddNewExerciseToWorkoutPlan(
	ath athlete.Athlete,
	wplan wp.WorkoutPlan,
	name string,
	metrics e.Metrics,
) (e.Exercise, error) {
	if wplan.AthleteID() != ath.AthleteID() {
		return e.Exercise{}, ErrUnauthorizedAccess
	}

	found, err := p.repo.FindWorkoutPlanByID(wplan)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return e.Exercise{}, errInternal
	}

	if !found {
		return e.Exercise{}, ErrWorkoutPlanNotFound
	}

	exercise, err := e.NewExercise(wplan.ID(), ath.AthleteID(), name, metrics)
	if err != nil {
		return e.Exercise{}, err
	}

	found, err = p.repo.FindExerciseByNameAndWorkoutPlanID(wplan, exercise)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return e.Exercise{}, errInternal
	}

	if found {
		return e.Exercise{}, ErrIdentialName
	}

	err = p.repo.StoreExercise(wplan, exercise, ath)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return e.Exercise{}, errInternal
	}

	return exercise, nil
}

func (p PlanningService) RemoveExerciseFromWorkoutPlan(
	ath athlete.Athlete,
	wplan wp.WorkoutPlan,
	exercise e.Exercise,
) error {
	if wplan.AthleteID() != ath.AthleteID() || exercise.AthleteID != ath.AthleteID() || exercise.WorkoutPlanID() != wplan.ID() {
		return ErrUnauthorizedAccess
	}

	found, err := p.repo.FindWorkoutPlanByID(wplan)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	if !found {
		return ErrWorkoutPlanNotFound
	}

	found, err = p.repo.FindExerciseByID(exercise)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	if !found {
		return ErrExerciseNotFound
	}

	err := p.repo.RemoveExercise(exercise)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	return nil
}
