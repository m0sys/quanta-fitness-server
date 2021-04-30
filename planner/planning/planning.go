package planning

import (
	"fmt"
	"log"

	"github.com/mhd53/quanta-fitness-server/athlete"
	"github.com/mhd53/quanta-fitness-server/planner/workoutplan"
)

const (
	errSlur = "planning"
)

var (
	errInternal      = fmt.Errorf("%s: internal error", errSlur)
	ErrIdentialTitle = fmt.Errorf("%s: WorkoutPlan with identical title already exists", errSlur)
)

type PlanningService struct {
	repo Repository
}

func NewPlanningService(repository Repository) PlanningService {
	return PlanningService{repo: repository}
}

func (p PlanningService) CreateNewWorkoutPlan(ath athlete.Athlete, title string) error {
	wplan, err := workoutplan.NewWorkoutPlan(title)
	if err != nil {
		return err
	}

	_, found, err := p.repo.FindWorkoutPlanByTitleAndAthleteID(title, ath)
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
