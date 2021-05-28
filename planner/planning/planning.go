package planning

import (
	"log"

	e "github.com/mhd53/quanta-fitness-server/planner/exercise"
	wp "github.com/mhd53/quanta-fitness-server/planner/workoutplan"
)

type PlanningService struct {
	repo Repository
}

func NewPlanningService(repository Repository) PlanningService {
	return PlanningService{
		repo: repository,
	}
}

// FIXME: assuming that Athlete IDs are valid - meaning that they are stored in manger.
// NOTE: I am ignoring security concerns for now.

func (p PlanningService) CreateNewWorkoutPlan(req CreateNewWorkoutPlanReq) (WorkoutPlanRes, error) {
	// TODO: check that aid is a valid Athlete ID - security layer.
	if err := ValidateCreateNewWorkoutPlanReq(req); err != nil {
		return WorkoutPlanRes{}, err
	}

	wplan, err := wp.NewWorkoutPlan(req.AthleteID, req.Title)
	if err != nil {
		return WorkoutPlanRes{}, err
	}

	found, err := p.repo.FindWorkoutPlanByTitleAndAthleteID(wplan)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return WorkoutPlanRes{}, errInternal
	}

	if found {
		return WorkoutPlanRes{}, ErrIdentialTitle
	}

	err = p.repo.StoreWorkoutPlan(wplan)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return WorkoutPlanRes{}, errInternal
	}

	res := mapWorkoutPlanToWorkoutPlanRes(wplan)
	return res, nil
}

func (p PlanningService) AddNewExerciseToWorkoutPlan(
	req AddNewExerciseToWorkoutPlanReq,
) (ExerciseRes, error) {
	if err := ValidateAddExerciseToWorkoutPlanReq(req); err != nil {
		return ExerciseRes{}, err
	}

	if err := p.validateWorkoutPlan(req.AthleteID, req.WorkoutPlanID); err != nil {
		return ExerciseRes{}, err
	}

	wplan, err := p.findWorkout(req.WorkoutPlanID)
	if err != nil {
		return ExerciseRes{}, err
	}

	pos, err := p.genPos(wplan)
	if err != nil {
		return ExerciseRes{}, err
	}

	metrics, err := e.NewMetrics(
		req.TargetRep,
		req.NumSets,
		req.Weight,
		req.RestDur,
	)
	if err != nil {
		return ExerciseRes{}, err
	}
	exercise, err := e.NewExercise(req.WorkoutPlanID, req.AthleteID, req.Name, metrics, pos)
	if err != nil {
		return ExerciseRes{}, err
	}

	found, err := p.repo.FindExerciseByNameAndWorkoutPlanID(req.WorkoutPlanID, req.Name)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return ExerciseRes{}, errInternal
	}

	if found {
		return ExerciseRes{}, ErrIdentialName
	}

	err = p.repo.StoreExercise(wplan, exercise)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return ExerciseRes{}, errInternal
	}

	res := mapExerciseToExerciseRes(exercise)
	return res, nil
}

func (p PlanningService) validateWorkoutPlan(aid, wpid string) error {
	wplan, err := p.findWorkout(wpid)
	if err != nil {
		return ErrWorkoutPlanNotFound
	}

	if !isAuthorizedWP(aid, wplan) {
		return ErrUnauthorizedAccess
	}

	return nil
}

func (p PlanningService) findWorkout(id string) (wp.WorkoutPlan, error) {
	wplan, found, err := p.repo.FindWorkoutPlanByID(id)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return wp.WorkoutPlan{}, errInternal
	}

	if !found {
		return wp.WorkoutPlan{}, ErrWorkoutPlanNotFound
	}

	return wplan, nil
}

func isAuthorizedWP(aid string, wplan wp.WorkoutPlan) bool {
	return wplan.AthleteID() == aid
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
	req RemoveExerciseFromWorkoutPlanReq,
) error {
	if err := ValidateRemoveExerciseFromWorkoutPlanReq(req); err != nil {
		return err
	}

	if err := p.validateWorkoutPlan(req.AthleteID, req.WorkoutPlanID); err != nil {
		return err
	}

	if err := p.validateExercise(req.AthleteID, req.WorkoutPlanID, req.ExerciseID); err != nil {
		return err
	}

	exercise, err := p.findExercise(req.ExerciseID)
	if err != nil {
		return err
	}

	err = p.repo.RemoveExercise(exercise)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return errInternal
	}

	return nil
}

func (p PlanningService) validateExercise(aid, wpid, eid string) error {
	exercise, err := p.findExercise(eid)
	if err != nil {
		return ErrExerciseNotFound
	}
	if !isAuthorizedE(aid, wpid, exercise) {
		return ErrUnauthorizedAccess
	}

	return nil
}

func (p PlanningService) findExercise(id string) (e.Exercise, error) {
	exercise, found, err := p.repo.FindExerciseByID(id)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return e.Exercise{}, errInternal
	}

	if !found {
		return e.Exercise{}, ErrExerciseNotFound
	}

	return exercise, nil
}

func isAuthorizedE(aid, wpid string, exercise e.Exercise) bool {
	return aid == exercise.AthleteID() && wpid == exercise.WorkoutPlanID()
}

func (p PlanningService) EditWorkoutPlanTitle(req EditWorkoutPlanTitleReq) error {
	if err := ValidateEditWorkoutPlanTitleReq(req); err != nil {
		return err
	}

	if err := p.validateWorkoutPlan(req.AthleteID, req.WorkoutPlanID); err != nil {
		return err
	}

	wplan, err := p.findWorkout(req.WorkoutPlanID)
	if err != nil {
		return err
	}

	err = wplan.EditTitle(req.Title)
	if err != nil {
		return err
	}

	found, err := p.repo.FindWorkoutPlanByTitleAndAthleteID(wplan)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return errInternal
	}

	if found {
		return ErrIdentialTitle
	}

	err = p.repo.UpdateWorkoutPlan(wplan)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return errInternal
	}

	return nil
}

func (p PlanningService) FetchWorkoutPlans(aid string) ([]WorkoutPlanRes, error) {
	var results []WorkoutPlanRes

	wplans, err := p.repo.FindAllWorkoutPlansForAthlete(aid)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return results, errInternal
	}

	for _, wplan := range wplans {
		res := mapWorkoutPlanToWorkoutPlanRes(wplan)
		results = append(results, res)
	}

	return results, nil
}

func (p PlanningService) FetchWorkoutPlanExercises(req FetchWorkoutPlanExercisesReq) ([]ExerciseRes, error) {
	var results []ExerciseRes

	if err := ValidateFetchWorkoutPlanExercisesReq(req); err != nil {
		return results, err
	}

	if err := p.validateWorkoutPlan(req.AthleteID, req.WorkoutPlanID); err != nil {
		return results, err
	}

	wplan, err := p.findWorkout(req.WorkoutPlanID)
	if err != nil {
		return results, err
	}

	exercises, err := p.repo.FindAllExercisesForWorkoutPlan(wplan)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return results, errInternal
	}

	for _, exercise := range exercises {
		res := mapExerciseToExerciseRes(exercise)
		results = append(results, res)
	}

	return results, nil
}

func (p PlanningService) EditExercise(req EditExerciseReq) error {
	if err := ValidateEditExerciseReq(req); err != nil {
		return err
	}

	if err := p.validateWorkoutPlan(req.AthleteID, req.WorkoutPlanID); err != nil {
		return err
	}

	if err := p.validateExercise(req.AthleteID, req.WorkoutPlanID, req.ExerciseID); err != nil {
		return err
	}

	exercise, err := p.findExercise(req.ExerciseID)
	if err != nil {
		return err
	}

	err = exercise.EditName(req.Name)
	if err != nil {
		return err
	}

	found, err := p.repo.FindExerciseByNameAndWorkoutPlanID(req.WorkoutPlanID, req.Name)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return errInternal
	}

	if found {
		return ErrIdentialName
	}

	err = exercise.EditTargetRep(req.TargetRep)
	if err != nil {
		return err
	}

	err = exercise.EditNumSets(req.NumSets)
	if err != nil {
		return err
	}

	err = exercise.EditWeight(req.Weight)
	if err != nil {
		return err
	}

	err = exercise.EditRestDur(req.RestDur)
	if err != nil {
		return err
	}

	err = p.repo.UpdateExercise(exercise)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return errInternal
	}

	return nil
}

func (p PlanningService) RemoveWorkoutPlan(req RemoveWorkoutPlanReq) error {
	if err := ValidateRemoveWorkoutPlanReq(req); err != nil {
		return err
	}

	if err := p.validateWorkoutPlan(req.AthleteID, req.WorkoutPlanID); err != nil {
		return err
	}

	wplan, err := p.findWorkout(req.WorkoutPlanID)
	if err != nil {
		return err
	}

	err = p.repo.RemoveWorkoutPlan(wplan)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return errInternal
	}

	return nil
}

// Mapper funcs.

func mapWorkoutPlanToWorkoutPlanRes(wplan wp.WorkoutPlan) WorkoutPlanRes {
	return WorkoutPlanRes{
		ID:        wplan.ID(),
		Title:     wplan.Title(),
		AthleteID: wplan.AthleteID(),
	}
}

func mapExerciseToExerciseRes(exercise e.Exercise) ExerciseRes {
	metrics := exercise.Metrics()
	return ExerciseRes{
		ID:            exercise.ID(),
		WorkoutPlanID: exercise.WorkoutPlanID(),
		AthleteID:     exercise.AthleteID(),
		Name:          exercise.Name(),
		TargetRep:     metrics.TargetRep(),
		NumSets:       metrics.NumSets(),
		Weight:        metrics.Weight(),
		RestDur:       metrics.RestDur(),
	}
}
