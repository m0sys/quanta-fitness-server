package translator

import (
	"log"

	p "github.com/mhd53/quanta-fitness-server/planner/planning"
	wp "github.com/mhd53/quanta-fitness-server/planner/workoutplan"
	elg "github.com/mhd53/quanta-fitness-server/tracker/exerciselog"
	t "github.com/mhd53/quanta-fitness-server/tracker/training"
	wl "github.com/mhd53/quanta-fitness-server/tracker/workoutlog"
)

type WorkoutTranslator struct {
	planRepo p.Repository
	logRepo  t.Repository
}

func NewWorkoutTranslator(planRepo p.Repository, logRepo t.Repository) WorkoutTranslator {
	return WorkoutTranslator{
		planRepo: planRepo,
		logRepo:  logRepo,
	}
}

/*
Preconditions:
	- `wplan` is a valid WorkoutPlan as defined in `PlanningService`.
*/
func (wt *WorkoutTranslator) ConvertWorkoutPlan(wplan wp.WorkoutPlan) (wl.WorkoutLog, error) {
	wlog := wl.NewWorkoutLog(wplan.AthleteID(), wplan.Title())
	exercises, err := wt.planRepo.FindAllExercisesForWorkoutPlan(wplan)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return wl.WorkoutLog{}, errInternal
	}

	if len(exercises) == 0 {
		return wl.WorkoutLog{}, ErrWorkoutPlanHasNoExercises
	}

	for _, e := range exercises {
		metrics := e.Metrics()
		lMetrics := elg.NewMetrics(metrics.TargetRep(), metrics.NumSets(), float64(metrics.Weight()), float64(metrics.RestDur()))
		elog := elg.NewExerciseLog(wlog.ID(), e.Name(), lMetrics, e.Pos())

		err = wt.logRepo.StoreExerciseLog(elog)
		if err != nil {
			log.Printf("%s: %s", errSlur, err.Error())
			return wl.WorkoutLog{}, errInternal
		}
	}

	err = wt.logRepo.StoreWorkoutLog(wlog)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return wl.WorkoutLog{}, errInternal
	}

	return wlog, nil
}
