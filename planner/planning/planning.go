package planning

import (
	"log"

	"github.com/mhd53/quanta-fitness-server/manager/athlete"
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

	found, err := p.repo.FindWorkoutPlanByTitleAndAthleteID(wplan, ath)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return wp.WorkoutPlan{}, errInternal
	}

	if found {
		return wp.WorkoutPlan{}, ErrIdentialTitle
	}

	err = p.repo.StoreWorkoutPlan(wplan, ath)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
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
	if err := p.validateWorkoutPlan(ath, wplan); err != nil {
		return e.Exercise{}, err
	}

	pos, err := p.genPos(wplan)
	if err != nil {
		return e.Exercise{}, err
	}

	exercise, err := e.NewExercise(wplan.ID(), ath.AthleteID(), name, metrics, pos)
	if err != nil {
		return e.Exercise{}, err
	}

	found, err := p.repo.FindExerciseByNameAndWorkoutPlanID(wplan, exercise)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return e.Exercise{}, errInternal
	}

	if found {
		return e.Exercise{}, ErrIdentialName
	}

	err = p.repo.StoreExercise(wplan, exercise, ath)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return e.Exercise{}, errInternal
	}

	return exercise, nil
}

func (p PlanningService) validateWorkoutPlan(ath athlete.Athlete, wplan wp.WorkoutPlan) error {
	if !isAuthorizedWP(ath, wplan) {
		return ErrUnauthorizedAccess
	}

	found, err := p.repo.FindWorkoutPlanByID(wplan)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return errInternal
	}

	if !found {
		return ErrWorkoutPlanNotFound
	}

	return nil
}

func isAuthorizedWP(ath athlete.Athlete, wplan wp.WorkoutPlan) bool {
	return wplan.AthleteID() == ath.AthleteID()
}

func (p PlanningService) genPos(wplan wp.WorkoutPlan) (int, error) {
	exercises, err := p.repo.FindAllExercisesForWorkoutPlan(wplan)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return -1, errInternal
	}

	length := len(exercises)

	if length == 0 {
		return 0, nil
	}

	return exercises[length-1].Pos() + 1, nil
}

func (p PlanningService) RemoveExerciseFromWorkoutPlan(
	ath athlete.Athlete,
	wplan wp.WorkoutPlan,
	exercise e.Exercise,
) error {
	if err := p.validateWorkoutPlan(ath, wplan); err != nil {
		return err
	}

	if err := p.validateExercise(ath, wplan, exercise); err != nil {
		return err
	}

	err := p.repo.RemoveExercise(exercise)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return errInternal
	}

	return nil
}

func (p PlanningService) validateExercise(ath athlete.Athlete, wplan wp.WorkoutPlan, exercise e.Exercise) error {
	if !isAuthorizedE(ath, wplan, exercise) {
		return ErrUnauthorizedAccess
	}

	found, err := p.repo.FindExerciseByID(exercise)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return errInternal
	}

	if !found {
		return ErrExerciseNotFound
	}
	return nil
}

func isAuthorizedE(ath athlete.Athlete, wplan wp.WorkoutPlan, exercise e.Exercise) bool {
	return ath.AthleteID() == exercise.AthleteID() && wplan.ID() == exercise.WorkoutPlanID()
}

func (p PlanningService) EditWorkoutPlanTitle(ath athlete.Athlete, wplan wp.WorkoutPlan, title string) error {
	if err := p.validateWorkoutPlan(ath, wplan); err != nil {
		return err
	}

	prevTitle := wplan.Title()

	err := wplan.EditTitle(title)
	if err != nil {
		wplan.EditTitle(prevTitle) //nolint
		return err
	}

	found, err := p.repo.FindWorkoutPlanByTitleAndAthleteID(wplan, ath)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		wplan.EditTitle(prevTitle) //nolint
		return errInternal
	}

	if found {
		wplan.EditTitle(prevTitle) //nolint
		return ErrIdentialTitle
	}

	err = p.repo.UpdateWorkoutPlan(wplan)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return errInternal
	}

	return nil
}

func (p PlanningService) FetchWorkoutPlans(ath athlete.Athlete) ([]wp.WorkoutPlan, error) {
	var wplans []wp.WorkoutPlan

	wplans, err := p.repo.FindAllWorkoutPlansForAthlete(ath)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return wplans, errInternal
	}

	return wplans, nil
}

func (p PlanningService) FetchWorkoutPlanExercises(ath athlete.Athlete, wplan wp.WorkoutPlan) ([]e.Exercise, error) {
	var exercises []e.Exercise

	if err := p.validateWorkoutPlan(ath, wplan); err != nil {
		return exercises, err
	}

	exercises, err := p.repo.FindAllExercisesForWorkoutPlan(wplan)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return exercises, errInternal
	}

	return exercises, nil
}

func (p PlanningService) EditExerciseName(ath athlete.Athlete, wplan wp.WorkoutPlan, exercise e.Exercise, name string) error {
	if err := p.validateWorkoutPlan(ath, wplan); err != nil {
		return err
	}

	if err := p.validateExercise(ath, wplan, exercise); err != nil {
		return err
	}

	prevName := exercise.Name()
	err := exercise.EditName(name)
	if err != nil {
		exercise.EditName(prevName) //nolint
		return err
	}

	found, err := p.repo.FindExerciseByNameAndWorkoutPlanID(wplan, exercise)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		exercise.EditName(prevName) //nolint
		return errInternal
	}

	if found {
		exercise.EditName(prevName) //nolint
		return ErrIdentialName
	}

	err = p.repo.UpdateExercise(exercise)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return errInternal
	}

	return nil
}

func (p PlanningService) EditExerciseTargetRep(ath athlete.Athlete, wplan wp.WorkoutPlan, exercise e.Exercise, targetRep int) error {
	if err := p.validateWorkoutPlan(ath, wplan); err != nil {
		return err
	}

	if err := p.validateExercise(ath, wplan, exercise); err != nil {
		return err
	}

	metrics := exercise.Metrics()
	prevTargetRep := metrics.TargetRep()
	err := exercise.EditTargetRep(targetRep)
	if err != nil {
		exercise.EditTargetRep(prevTargetRep) //nolint
		return err
	}

	err = p.repo.UpdateExercise(exercise)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		exercise.EditTargetRep(prevTargetRep) //nolint
		return errInternal
	}

	return nil
}

func (p PlanningService) EditExerciseNumSets(ath athlete.Athlete, wplan wp.WorkoutPlan, exercise e.Exercise, numSets int) error {
	if err := p.validateWorkoutPlan(ath, wplan); err != nil {
		return err
	}

	if err := p.validateExercise(ath, wplan, exercise); err != nil {
		return err
	}

	metrics := exercise.Metrics()
	prevNumSets := metrics.NumSets()
	err := exercise.EditNumSets(numSets)
	if err != nil {
		exercise.EditNumSets(prevNumSets) //nolint
		return err
	}

	err = p.repo.UpdateExercise(exercise)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		exercise.EditNumSets(prevNumSets) //nolint
		return errInternal
	}

	return nil
}

func (p PlanningService) EditExerciseWeight(ath athlete.Athlete, wplan wp.WorkoutPlan, exercise e.Exercise, weight float64) error {
	if err := p.validateWorkoutPlan(ath, wplan); err != nil {
		return err
	}

	if err := p.validateExercise(ath, wplan, exercise); err != nil {
		return err
	}

	metrics := exercise.Metrics()
	prevWeight := float64(metrics.Weight())
	err := exercise.EditWeight(weight)
	if err != nil {
		exercise.EditWeight(prevWeight) //nolint
		return err
	}

	err = p.repo.UpdateExercise(exercise)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		exercise.EditWeight(prevWeight) //nolint
		return errInternal
	}

	return nil
}

func (p PlanningService) EditExerciseRestDur(ath athlete.Athlete, wplan wp.WorkoutPlan, exercise e.Exercise, restDur float64) error {
	if err := p.validateWorkoutPlan(ath, wplan); err != nil {
		return err
	}

	if err := p.validateExercise(ath, wplan, exercise); err != nil {
		return err
	}

	metrics := exercise.Metrics()
	prevRestDur := float64(metrics.RestDur())
	err := exercise.EditRestDur(restDur)
	if err != nil {
		exercise.EditRestDur(prevRestDur) //nolint
		return err
	}

	err = p.repo.UpdateExercise(exercise)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		exercise.EditRestDur(prevRestDur) //nolint
		return errInternal
	}

	return nil
}

func (p PlanningService) RemoveWorkoutPlan(ath athlete.Athlete, wplan wp.WorkoutPlan) error {
	if err := p.validateWorkoutPlan(ath, wplan); err != nil {
		return err
	}

	err := p.repo.RemoveWorkoutPlan(wplan)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return errInternal
	}

	return nil
}
