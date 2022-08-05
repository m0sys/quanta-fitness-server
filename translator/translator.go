package translator

import (
	"log"

	p "github.com/m0sys/quanta-fitness-server/planner/planning"
	elg "github.com/m0sys/quanta-fitness-server/tracker/exerciselog"
	t "github.com/m0sys/quanta-fitness-server/tracker/training"
	wl "github.com/m0sys/quanta-fitness-server/tracker/workoutlog"
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

func (wt *WorkoutTranslator) ConvertWorkoutPlan(res p.WorkoutPlanRes) (wl.WorkoutLog, error) {
	wlog := wl.NewWorkoutLog(res.AthleteID, res.Title)
	wplan, found, err := wt.planRepo.FindWorkoutPlanByID(res.ID)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return wl.WorkoutLog{}, errInternal
	}

	if !found {
		log.Printf("%s: %s", errSlur, "WorkoutPlan not found")
		return wl.WorkoutLog{}, errInternal
	}

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
