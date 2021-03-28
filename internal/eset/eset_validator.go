package eset

import (
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type EsetValidator interface {
	ValidateAddEsetToExercise(actualRC int, dur, restDur float32) error
	ValidateUpdateEset(updates entity.EsetUpdate) error
}
