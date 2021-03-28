package eset

import (
	// "errors"
	// "log"

	// store "github.com/mhd53/quanta-fitness-server/internal/datastore/esetstore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type EsetService interface {
	AddEsetToExercise(uname string, eid int64, actualRC int, dur, restDur float32) (entity.Eset, error)
	UpdateEset(esid int64, uname string, updates entity.EsetUpdate) error
	DeleteEset(esid int64, uname string) error
	GetEset(esid int64, uname string) (entity.Eset, error)
	GetEsetsForExercise(eid int64, uname string) ([]entity.Eset, error)
}
