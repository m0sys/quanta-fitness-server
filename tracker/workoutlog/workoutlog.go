package workoutlog

import (
	"time"

	"github.com/mhd53/quanta-fitness-server/pkg/uuid"
)

type WorkoutLog struct {
	uuid      string
	aid       string
	title     string
	date      time.Time
	completed bool
}

// This constructor should only be used through WorkoutPlanToWorkoutLogTranslator.
func NewWorkoutLog(aid, title string) WorkoutLog {
	return WorkoutLog{
		uuid:      uuid.GenerateUUID(),
		aid:       aid,
		title:     title,
		date:      time.Now(),
		completed: false,
	}
}

// This should only be used to restore data from persistence layer.
func RestoreWorkoutLog(id, aid, title string, date time.Time, completed bool) WorkoutLog {
	return WorkoutLog{
		uuid:      id,
		aid:       aid,
		title:     title,
		date:      date,
		completed: completed,
	}
}

func (w *WorkoutLog) ID() string {
	return w.uuid
}

func (w *WorkoutLog) AthleteID() string {
	return w.aid
}

func (w *WorkoutLog) Title() string {
	return w.title
}

func (w *WorkoutLog) Date() time.Time {
	return w.date
}

func (w *WorkoutLog) Completed() bool {
	return w.completed
}

func (w *WorkoutLog) Complete() {
	w.completed = true
}
