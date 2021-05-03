package training

import (
	"log"

	"github.com/mhd53/quanta-fitness-server/manager/athlete"
	el "github.com/mhd53/quanta-fitness-server/tracker/exerciselog"
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

	if !isAuthorizedWL(ath, wlog) {
		return elogs, ErrUnauthorizedAccess
	}

	found, err := t.repo.FindWorkoutLogByID(wlog)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return elogs, errInternal
	}

	if !found {
		return elogs, ErrWorkoutLogNotFound

	}

	elogs, err = t.repo.FindAllExerciseLogsForWorkoutLog(wlog)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return elogs, errInternal
	}

	return elogs, nil
}

func isAuthorizedWL(ath athlete.Athlete, wlog wl.WorkoutLog) bool {
	return wlog.AthleteID() == ath.AthleteID()
}
