package training

import (
	"log"

	"github.com/mhd53/quanta-fitness-server/manager/athlete"
	el "github.com/mhd53/quanta-fitness-server/tracker/exerciselog"
	sl "github.com/mhd53/quanta-fitness-server/tracker/setlog"
	wl "github.com/mhd53/quanta-fitness-server/tracker/workoutlog"
)

type TrainingService struct {
	repo Repository
}

func NewTrainingService(repository Repository) TrainingService {
	return TrainingService{repo: repository}
}

func (t TrainingService) FetchWorkoutLogs(ath athlete.Athlete) ([]wl.WorkoutLog, error) {
	var wlogs []wl.WorkoutLog

	wlogs, err := t.repo.FindAllWorkoutLogsForAthlete(ath)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return wlogs, errInternal
	}

	return wlogs, nil
}

func (t TrainingService) FetchWorkoutLogExerciseLogs(
	ath athlete.Athlete,
	wlog wl.WorkoutLog,
) ([]el.ExerciseLog, error) {
	var elogs []el.ExerciseLog

	if err := t.validateWlog(ath, wlog); err != nil {
		return elogs, err
	}

	elogs, err := t.repo.FindAllExerciseLogsForWorkoutLog(wlog)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return elogs, errInternal
	}

	return elogs, nil
}

func (t TrainingService) validateWlog(ath athlete.Athlete, wlog wl.WorkoutLog) error {
	if !isAuthorizedWL(ath, wlog) {
		return ErrUnauthorizedAccess
	}

	found, err := t.repo.FindWorkoutLogByID(wlog)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	if !found {
		return ErrWorkoutLogNotFound

	}
	return nil
}

func isAuthorizedWL(ath athlete.Athlete, wlog wl.WorkoutLog) bool {
	return wlog.AthleteID() == ath.AthleteID()
}

func (t TrainingService) AddSetLogToExerciseLog(ath athlete.Athlete, wlog wl.WorkoutLog, elog el.ExerciseLog, metrics sl.Metrics) error {
	if err := t.validateWlog(ath, wlog); err != nil {
		return err
	}

	if err := t.validateElog(ath, wlog, elog); err != nil {
		return err
	}

	slogs, err := t.repo.FindAllSetLogsForExerciseLog(elog)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	eMetrics := elog.Metrics()
	if len(slogs) == eMetrics.NumSets() {
		return ErrCannotExceedNumSets

	}

	slog := sl.NewSetLog(elog.ID(), metrics)
	err = t.repo.StoreSetLog(slog)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	return nil
}

func isAuthorizedEL(ath athlete.Athlete, wlog wl.WorkoutLog, elog el.ExerciseLog) bool {
	return ath.AthleteID() == wlog.AthleteID() && elog.WorkoutLogID() == wlog.ID()
}

func (t TrainingService) validateElog(ath athlete.Athlete, wlog wl.WorkoutLog, elog el.ExerciseLog) error {
	if !isAuthorizedEL(ath, wlog, elog) {
		return ErrUnauthorizedAccess
	}

	found, err := t.repo.FindExerciseLogByID(elog)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	if !found {
		return ErrExerciseLogNotFound

	}

	return nil

}

func (t TrainingService) MoveToNextExerciseLog(ath athlete.Athlete, wlog wl.WorkoutLog) (wl.WorkoutLog, el.ExerciseLog, error) {
	if err := t.validateWlog(ath, wlog); err != nil {
		return wlog, el.ExerciseLog{}, err
	}

	if wlog.Completed() {
		return wlog, el.ExerciseLog{}, ErrWorkoutLogAlreadyCompleted
	}

	elogs, err := t.repo.FindAllExerciseLogsForWorkoutLog(wlog)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return wlog, el.ExerciseLog{}, errInternal
	}

	prevState := wl.RestoreWorkoutLog(wlog.ID(), wlog.AthleteID(), wlog.Title(), wlog.Date(), wlog.CurrentPos(), wlog.Completed())
	currPos := wlog.CurrentPos()
	wlog.NextPos()
	if wlog.CurrentPos() == len(elogs) {
		wlog.Complete()
	}

	err = t.repo.UpdateWorkoutLog(wlog)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return prevState, el.ExerciseLog{}, errInternal
	}

	return wlog, elogs[currPos], nil
}

func (t TrainingService) FetchSetLogsForExerciseLog(ath athlete.Athlete, wlog wl.WorkoutLog, elog el.ExerciseLog) ([]sl.SetLog, error) {
	var slogs []sl.SetLog
	if err := t.validateWlog(ath, wlog); err != nil {
		return slogs, err
	}

	if err := t.validateElog(ath, wlog, elog); err != nil {
		return slogs, err
	}

	slogs, err := t.repo.FindAllSetLogsForExerciseLog(elog)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return slogs, errInternal
	}

	return slogs, nil
}
