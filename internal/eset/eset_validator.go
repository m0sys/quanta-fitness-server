package eset

import (
	"errors"

	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type EsetValidator interface {
	ValidateAddEsetToExercise(actualRC int, dur, restDur float32) error
	ValidateUpdateEset(updates entity.EsetUpdate) error
}

type validator struct{}

func NewEsetValidator() EsetValidator {
	return &validator{}
}

func (*validator) ValidateAddEsetToExercise(actualRC int, dur, restDur float32) error {
	return validateFields(actualRC, dur, restDur)
}

func (*validator) ValidateUpdateEset(updates entity.EsetUpdate) error {
	return validateFields(updates.SMetric.ActualRepCount, updates.SMetric.Duration, updates.SMetric.RestTimeDuration)
}

// Helper funcs.

func validateFields(arc int, dur, restDur float32) error {
	err := validateActualRC(arc)

	if err != nil {
		return err
	}

	err2 := validateDur(dur)

	if err2 != nil {
		return err2
	}

	err3 := validateRestDur(restDur)

	if err3 != nil {
		return err3
	}

	return nil
}

func validateActualRC(actualRC int) error {
	if isNegativeInt(actualRC) {
		return errors.New("Sign Error: Actual rep count must be positive or zero!")
	}
	return nil
}

func isNegativeInt(i int) bool {
	return i < 0
}

func validateDur(dur float32) error {
	if isNegativeFloat(dur) {
		return errors.New("Sign Error: Duration must be positive or zero!")
	}
	return nil
}

func isNegativeFloat(f float32) bool {
	return f < 0.0
}

func validateRestDur(restDur float32) error {
	if isNegativeFloat(restDur) {
		return errors.New("Sign Error: Rest duration must be positive or zero!")
	}
	return nil
}
