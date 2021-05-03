package exercise

import (
	"errors"
	"math"

	"github.com/mhd53/quanta-fitness-server/pkg/uuid"
	"github.com/mhd53/quanta-fitness-server/units"
)

type Exercise struct {
	uuid    string
	wpid    string
	aid     string
	name    string
	metrics Metrics
	pos     int
}

type Metrics struct {
	targetRep int
	numSets   int
	weight    units.Kilogram
	restDur   units.Second
}

func NewExercise(wpid, aid, name string, metrics Metrics, pos int) (Exercise, error) {
	if err := validateName(name); err != nil {
		return Exercise{}, err
	}

	if err := validatePos(pos); err != nil {
		return Exercise{}, err
	}

	return Exercise{
		uuid:    uuid.GenerateUUID(),
		wpid:    wpid,
		aid:     aid,
		name:    name,
		metrics: metrics,
		pos:     pos,
	}, nil
}

// FIXME: find alternative solution for id checking...
func RestoreExercise(id, wpid, aid, name string, metrics Metrics, pos int) (Exercise, error) {
	if err := validateName(name); err != nil {
		return Exercise{}, err
	}

	if err := validatePos(pos); err != nil {
		return Exercise{}, err
	}

	return Exercise{
		uuid:    id,
		wpid:    wpid,
		aid:     aid,
		name:    name,
		metrics: metrics,
		pos:     pos,
	}, nil
}

func (e *Exercise) EditName(name string) error {
	if err := validateName(name); err != nil {
		return err
	}

	e.name = name
	return nil
}

func (e *Exercise) EditTargetRep(targetRep int) error {
	newMetrics, err := NewMetrics(
		targetRep,
		e.metrics.NumSets(),
		float64(e.metrics.Weight()),
		float64(e.metrics.RestDur()),
	)
	if err != nil {
		return err
	}

	e.metrics = newMetrics
	return nil
}

func (e *Exercise) EditNumSets(numSets int) error {
	newMetrics, err := NewMetrics(
		e.metrics.TargetRep(),
		numSets,
		float64(e.metrics.Weight()),
		float64(e.metrics.RestDur()),
	)
	if err != nil {
		return err
	}

	e.metrics = newMetrics
	return nil
}

func (e *Exercise) EditWeight(weight float64) error {
	newMetrics, err := NewMetrics(
		e.metrics.TargetRep(),
		e.metrics.NumSets(),
		weight,
		float64(e.metrics.RestDur()),
	)
	if err != nil {
		return err
	}

	e.metrics = newMetrics
	return nil
}

func (e *Exercise) EditRestDur(restDur float64) error {
	newMetrics, err := NewMetrics(
		e.metrics.TargetRep(),
		e.metrics.NumSets(),
		float64(e.metrics.Weight()),
		restDur,
	)
	if err != nil {
		return err
	}

	e.metrics = newMetrics
	return nil
}

func (e *Exercise) ID() string {
	return e.uuid
}

func (e *Exercise) Name() string {
	return e.name
}

func (e *Exercise) Metrics() Metrics {
	return e.metrics
}

func (e *Exercise) AthleteID() string {
	return e.aid
}

func (e *Exercise) WorkoutPlanID() string {
	return e.wpid
}

func (e *Exercise) Pos() int {
	return e.pos
}

func NewMetrics(targetRep, numSets int, weight, restDur float64) (Metrics, error) {
	if err := validateMetrics(targetRep, numSets, weight, restDur); err != nil {
		return Metrics{}, err
	}

	return Metrics{
		targetRep: targetRep,
		numSets:   numSets,
		weight:    units.Kilogram(roundToTwoDecimalPlaces(weight)),
		restDur:   units.Second(roundToTwoDecimalPlaces(restDur)),
	}, nil
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

// Helper funcs.

func validateMetrics(targetRep, numSets int, weight, restDur float64) error {
	if err := validateWeight(weight); err != nil {
		return err
	}

	if err := validateRestDur(restDur); err != nil {
		return err
	}

	if err := validateTargetRep(targetRep); err != nil {
		return err
	}

	if err := validateNumSets(numSets); err != nil {
		return err
	}

	return nil
}

var (
	ErrInvalidName      = errors.New("name must be less than 76 characters")
	ErrInvalidTargetRep = errors.New("target rep must be a positive number")
	ErrInvalidNumSets   = errors.New("num sets must be a positive number")
	ErrInvalidWeight    = errors.New("weight must be a positive number")
	ErrInvalidRestDur   = errors.New("rest duration must be a positive number")
	ErrInvalidPos       = errors.New("position must be a positive number")
)

func validateName(name string) error {
	if len(name) > 75 {
		return ErrInvalidName
	}

	return nil
}

func validatePos(pos int) error {
	if pos < 0 {
		return ErrInvalidPos
	}
	return nil
}

func validateTargetRep(targetRep int) error {
	if targetRep < 0 {
		return ErrInvalidTargetRep
	}
	return nil
}

func validateNumSets(numSets int) error {
	if numSets < 0 {
		return ErrInvalidNumSets
	}
	return nil
}

func validateWeight(weight float64) error {
	if weight < 0 {
		return ErrInvalidWeight
	}
	return nil
}

func validateRestDur(restDur float64) error {
	if restDur < 0 {
		return ErrInvalidRestDur
	}
	return nil
}

func roundToTwoDecimalPlaces(num float64) float64 {
	return math.Round(num*100) / 100
}
