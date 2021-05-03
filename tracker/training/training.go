package training

import (
	"log"

	"github.com/mhd53/quanta-fitness-server/manager/athlete"
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
