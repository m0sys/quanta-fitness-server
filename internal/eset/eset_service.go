package eset

import (
	"errors"
	"fmt"
	"log"

	store "github.com/m0sys/quanta-fitness-server/internal/datastore/esetstore"
	"github.com/m0sys/quanta-fitness-server/internal/entity"
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
		logErr(err)
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
			ActualRepCount:   actualRC,
			Duration:         dur,
			RestTimeDuration: restDur,
		},
	})
	if err3 != nil {
		return entity.Eset{}, formatErr(err3)
	}

	return added, nil

}

func (*service) UpdateEset(esid int64, uname string, updates entity.EsetUpdate) error {
	authorized, err := auth.AuthorizeEsetAccess(uname, esid)

	if err != nil {
		logErr(err)
		return internalErr
	}

	if !authorized {
		return deniedErr
	}

	err2 := val.ValidateUpdateEset(updates)
	if err2 != nil {
		return err2
	}

	err3 := sess.Update(esid, updates)
	if err3 != nil {
		return formatErr(err3)
	}

	return nil
}

func (*service) DeleteEset(esid int64, uname string) error {
	authorized, err := auth.AuthorizeEsetAccess(uname, esid)

	if err != nil {
		logErr(err)
		return internalErr
	}

	if !authorized {
		return deniedErr
	}

	err2 := sess.Delete(esid)

	if err2 != nil {
		return formatErr(err2)
	}

	return nil
}

func (*service) GetEset(esid int64, uname string) (entity.Eset, error) {
	authorized, err := auth.AuthorizeEsetAccess(uname, esid)

	if err != nil {
		logErr(err)
		return entity.Eset{}, internalErr
	}

	if !authorized {
		return entity.Eset{}, deniedErr
	}

	got, _, err2 := sess.FindEsetById(esid)

	if err2 != nil {
		return entity.Eset{}, formatErr(err2)
	}

	return got, nil

}

func (*service) GetEsetsForExercise(eid int64, uname string) ([]entity.Eset, error) {
	var esets []entity.Eset

	authorized, err := auth.AuthorizeExerciseAccess(uname, eid)

	if err != nil {
		return esets, err
	}

	if !authorized {
		return esets, deniedErr
	}

	esets, err2 := sess.FindAllEsetByEID(eid)

	if err2 != nil {
		return esets, formatErr(err2)
	}

	return esets, nil
}

func logErr(err error) {
	log.Printf("%s: %s", "eset_service", err.Error())
}

func formatErr(err error) error {
	return fmt.Errorf("%s: couldn't access db: %w", "eset_service", err)
}
