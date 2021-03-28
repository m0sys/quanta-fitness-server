package eset

import (
	"errors"
	"log"

	store "github.com/mhd53/quanta-fitness-server/internal/datastore/esetstore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type EsetService interface {
	AddEsetToExercise(uname string, eid int64, actualRC int, dur, restDur float32) (entity.Eset, error)
	UpdateEset(esid int64, uname string, updates entity.EsetUpdate) error
	DeleteEset(esid int64, uname string) error
	GetEset(esid int64, uname string) (entity.Eset, error)
	GetEsetsForExercise(eid int64, uname string) ([]entity.Eset, error)
}

var (
	sess        store.EsetStore
	auth        EsetAuth
	val         EsetValidator
	deniedErr   = errors.New("Access Denied!")
	internalErr = errors.New("Internal Error!")
	notImplErr  = errors.New("Not Implemented!")
)

type service struct{}

func NewEsetService(
	esstore store.EsetStore,
	authorizer EsetAuth,
	validator EsetValidator) EsetService {
	sess = esstore
	auth = authorizer
	val = validator
	return &service{}
}

func (*service) AddEsetToExercise(uname string, eid int64, actualRC int, dur, restDur float32) (entity.Eset, error) {
	authorized, err := auth.AuthorizeExerciseAccess(uname, eid)

	if err != nil {
		log.Panic(err)
		return entity.Eset{}, internalErr
	}

	if !authorized {
		return entity.Eset{}, deniedErr
	}

	err2 := val.ValidateAddEsetToExercise(actualRC, dur, restDur)

	if err2 != nil {
		return entity.Eset{}, err2
	}

	added, err3 := sess.Save(entity.BaseEset{
		Username: uname,
		EID:      eid,
		SMetric: entity.SMetric{
			ActualRepCount:    actualRC,
			Duraction:         dur,
			RestTimeDuraction: restDur,
		},
	})
	if err3 != nil {
		log.Panic(err3)
		return entity.Eset{}, internalErr
	}

	return added, nil

}

func (*service) UpdateEset(esid int64, uname string, updates entity.EsetUpdate) error {
	return notImplErr
}

func (*service) DeleteEset(esid int64, uname string) error {
	return notImplErr
}

func (*service) GetEset(esid int64, uname string) (entity.Eset, error) {
	return entity.Eset{}, notImplErr
}

func (*service) GetEsetsForExercise(eid int64, uname string) ([]entity.Eset, error) {
	var esets []entity.Eset
	return esets, notImplErr
}
