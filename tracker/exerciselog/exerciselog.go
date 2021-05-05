package exerciselog

import (
	"github.com/mhd53/quanta-fitness-server/pkg/uuid"
	"github.com/mhd53/quanta-fitness-server/units"
)

type ExerciseLog struct {
	uuid      string
	wlid      string
	name      string
	metrics   Metrics
	completed bool
	pos       int
}

type Metrics struct {
	targetRep int
	numSets   int
	weight    units.Kilogram
	restDur   units.Second
}

// This constructor should only be used through ExercisePlanToExerciseLogTranslator.
func NewExerciseLog(wlid, name string, metrics Metrics, pos int) ExerciseLog {
	return ExerciseLog{
		uuid:      uuid.GenerateUUID(),
		wlid:      wlid,
		name:      name,
		metrics:   metrics,
		completed: false,
		pos:       pos,
	}
}

// This should only be used to restore data from persistence layer.
func RestoreExerciseLog(id, wlid, name string, completed bool, metrics Metrics, pos int) ExerciseLog {
	return ExerciseLog{
		uuid:      id,
		wlid:      wlid,
		name:      name,
		metrics:   metrics,
		completed: completed,
		pos:       pos,
	}
}

func (e *ExerciseLog) ID() string {
	return e.uuid
}

func (e *ExerciseLog) WorkoutLogID() string {
	return e.wlid
}

func (e *ExerciseLog) Name() string {
	return e.name
}

func (e *ExerciseLog) Metrics() Metrics {
	return e.metrics
}

func (e *ExerciseLog) Completed() bool {
	return e.completed
}

func (e *ExerciseLog) Complete() {
	e.completed = true
}

func (e *ExerciseLog) Pos() int {
	return e.pos
}

func NewMetrics(targetRep, numSets int, weight, restDur float64) Metrics {
	return Metrics{
		targetRep: targetRep,
		numSets:   numSets,
		weight:    units.Kilogram(weight),
		restDur:   units.Second(restDur),
	}
}

func (m *Metrics) TargetRep() int {
	return m.targetRep
}

func (m *Metrics) NumSets() int {
	return m.numSets
}

func (m *Metrics) Weight() units.Kilogram {
	return m.weight
}

func (m *Metrics) RestDur() units.Second {
	return m.restDur
}
