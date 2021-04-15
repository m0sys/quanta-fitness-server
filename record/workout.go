package record

import "time"

type Workout struct {
	AID       int64
	ID        int64
	Title     string
	Date      time.Time
	Exercises []Exercise
	State     string
}

type WorkoutBuilder interface {
	AddNewExercise(e Exercise)
}

type WorkoutSessionManger interface {
	Start()
	MoveToNextExercise() func() bool
	Complete()
}

func (w *Workout) AddNewExercise(e Exercise) {
	w.exercises = append(w.exercises, e)
}

func (w *Workout) GetExercises() []Exercise {
	return w.exercises
}

func (w *Workout) GetState() string {
	return w.state
}

func (w *Workout) Start() {
	w.state = "inprogress"
}

func (w *Workout) MoveToNextExercise() func() bool {
	lastIdx := 0

	return func() {
		if lastIdx+1 < len(w.exercises) {
			lastIdx += 1
			return true
		}
		return false
	}

}

func (w *Workout) Complete() {
	w.state = "complete"
}

// WorkoutBuilder.

type workoutBuilder struct{}

func NewWorkoutBuilder() WorkoutBuilder {
	return &workoutBuilder{}
}

func (b *workoutBuilder) CreateNewWorkout(title string, date time.Time) Workout {
	return Workout{
		Title: title,
		Date:  date,
		State: "incomplete",
	}
}

func (b *workoutBuilder) AddNewExercise(w Workout, e Exercise) Workout {
	var exercises []Exercise
	copy(exercises, w.Exercises)

	var sets []Set
	copy(sets, e.Sets)

	exercise = Exercise{
		Name:      e.Name,
		RestTime:  e.RestTime,
		Distance:  e.Distance,
		TargetRep: e.TargetRep,
		Sets:      sets,
	}
	exercises = append(exercises, e)

	return Workout{
		Title:     title,
		Date:      date,
		Exercises: exercises,
	}

}

func (b *workoutBuilder) EditWorkout(w Workout, title string, data time.Time) Workout {
	var exercises []Exercise
	copy(exercises, w.Exercises)

	return Workout{
		Title:     title,
		Date:      date,
		Exercises: exercises,
	}
}

// WorkoutSessionManger.

type workoutSessionManger struct{}

func NewWorkoutSessionManger() WorkoutSessionManger {
	return &workoutSessionManger
}

func (m *workoutSessionManger) Start(w *Workout) {
	w.State = "inprogress"
}

func (m *workoutSessionManger) MoveToNextExercise() func() bool {
	return false

}

func (m *workoutSessionManger) Complete() {
	return
}
