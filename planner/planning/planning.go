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
	if !isAuthorizedWP(ath, wplan) {
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

func isAuthorizedWP(ath athlete.Athlete, wplan wp.WorkoutPlan) bool {
	return wplan.AthleteID() == ath.AthleteID()
}

func (p PlanningService) RemoveExerciseFromWorkoutPlan(
	ath athlete.Athlete,
	wplan wp.WorkoutPlan,
	exercise e.Exercise,
) error {
	if !isAuthorizedWP(ath, wplan) || !isAuthorizedE(ath, wplan, exercise) {
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

	err = p.repo.RemoveExercise(exercise)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	return nil
}

func isAuthorizedE(ath athlete.Athlete, wplan wp.WorkoutPlan, exercise e.Exercise) bool {
	return ath.AthleteID() == exercise.AthleteID() && wplan.ID() == exercise.WorkoutPlanID()
}

func (p PlanningService) EditWorkoutPlanTitle(ath athlete.Athlete, wplan wp.WorkoutPlan, title string) error {

	if !isAuthorizedWP(ath, wplan) {
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

	err = wplan.EditTitle(title)
	err = p.repo.UpdateWorkoutPlan(wplan)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	return nil
}

func (p PlanningService) FetchWorkoutPlans(ath athlete.Athlete) ([]wp.WorkoutPlan, error) {
	var wplans []wp.WorkoutPlan

	wplans, err := p.repo.FindAllWorkoutPlansForAthlete(ath)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return wplans, errInternal
	}

	return wplans, nil
}

func (p PlanningService) FetchWorkoutPlanExercises(ath athlete.Athlete, wplan wp.WorkoutPlan) ([]e.Exercise, error) {
	var exercises []e.Exercise

	if !isAuthorizedWP(ath, wplan) {
		return exercises, ErrUnauthorizedAccess
	}

	found, err := p.repo.FindWorkoutPlanByID(wplan)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return exercises, errInternal
	}

	if !found {
		return exercises, ErrWorkoutPlanNotFound
	}

	exercises, err = p.repo.FindAllExercisesForWorkoutPlan(wplan)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return exercises, errInternal
	}

	return exercises, nil
}

func (p PlanningService) EditExerciseName(ath athlete.Athlete, wplan wp.WorkoutPlan, exercise e.Exercise, name string) error {
	if !isAuthorizedWP(ath, wplan) || !isAuthorizedE(ath, wplan, exercise) {
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

	err = exercise.EditName(name)
	if err != nil {
		return err
	}

	err = p.repo.UpdateExercise(exercise)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	return nil
}

func (p PlanningService) EditExerciseTargetRep(ath athlete.Athlete, wplan wp.WorkoutPlan, exercise e.Exercise, targetRep int) error {
	if !isAuthorizedWP(ath, wplan) || !isAuthorizedE(ath, wplan, exercise) {
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

	err = exercise.EditTargetRep(targetRep)
	if err != nil {
		return err
	}

	err = p.repo.UpdateExercise(exercise)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	return nil
}

func (p PlanningService) EditExerciseNumSets(ath athlete.Athlete, wplan wp.WorkoutPlan, exercise e.Exercise, numSets int) error {
	if !isAuthorizedWP(ath, wplan) || !isAuthorizedE(ath, wplan, exercise) {
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

	err = exercise.EditNumSets(numSets)
	if err != nil {
		return err
	}

	err = p.repo.UpdateExercise(exercise)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	return nil
}

func (p PlanningService) EditExerciseWeight(ath athlete.Athlete, wplan wp.WorkoutPlan, exercise e.Exercise, weight float64) error {
	if !isAuthorizedWP(ath, wplan) || !isAuthorizedE(ath, wplan, exercise) {
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

	err = exercise.EditWeight(weight)
	if err != nil {
		return err
	}

	err = p.repo.UpdateExercise(exercise)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	return nil
}

func (p PlanningService) EditExerciseRestDur(ath athlete.Athlete, wplan wp.WorkoutPlan, exercise e.Exercise, restDur float64) error {
	if !isAuthorizedWP(ath, wplan) || !isAuthorizedE(ath, wplan, exercise) {
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

	err = exercise.EditRestDur(restDur)
	if err != nil {
		return err
	}

	err = p.repo.UpdateExercise(exercise)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	return nil
}
