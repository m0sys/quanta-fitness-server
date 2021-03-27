package exercise

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

var (
	validate = validator.New()
)

type ExerciseValidator interface {
	ValidateCreateExercise(name string) error
	ValidateUpdateExercise(updates entity.ExerciseUpdate) error
}

type evalidator struct{}

func NewExerciseValidator() ExerciseValidator {
	return &evalidator{}
}

func (*evalidator) ValidateCreateExercise(name string) error {
	return validateName(name)

}

func (*evalidator) ValidateUpdateExercise(updates entity.ExerciseUpdate) error {
	err := validateName(updates.Name)

	if err != nil {
		return err
	}

	err2 := validateWeight(updates.Metrics.Weight)
	if err2 != nil {
		return err2
	}

	err3 := validateRestTime(updates.Metrics.RestTime)
	if err3 != nil {
		return err3
	}

	err4 := validateTargetRep(updates.Metrics.TargetRep)
	if err4 != nil {
		return err4
	}

	err5 := validateNumSets(updates.Metrics.NumSets)
	if err5 != nil {
		return err5
	}

	return nil
}

// Helper functions.

func validateName(name string) error {
	err := checkNameLength(name)
	if err != nil {
		return err
	}

	return nil

}

func checkNameLength(name string) error {
	if validate.Var(name, "required,max=38") != nil {
		return errors.New("Name must be less than 38 characters!")
	}
	return nil
}

func validateWeight(weight float32) error {
	if isNegativeFloat(weight) {
		return errors.New("Sign Error: Weight must be positive or zero!")
	}
	return nil
}

func isNegativeFloat(f float32) bool {
	return f < 0.0
}

func validateRestTime(restTime float32) error {
	if isNegativeFloat(restTime) {
		return errors.New("Sign Error: Rest time must be positive or zero!")
	}
	return nil
}

func validateTargetRep(targetRep int) error {
	if isNegativeInt(targetRep) {
		return errors.New("Sign Error: Target rep must be positive or zero!")
	}
	return nil
}

func isNegativeInt(i int) bool {
	return i < 0
}

func validateNumSets(numSets int) error {
	if isNegativeInt(numSets) {
		return errors.New("Sign Error: Number of sets must be positive or zero!")
	}
	return nil
}
