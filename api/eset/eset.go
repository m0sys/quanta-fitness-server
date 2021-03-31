package eset

import (
	"errors"
	"log"

	esstore "github.com/mhd53/quanta-fitness-server/internal/datastore/esetstore"
	estore "github.com/mhd53/quanta-fitness-server/internal/datastore/exercisestore"
	ustore "github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
	wstore "github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
	serv "github.com/mhd53/quanta-fitness-server/internal/eset"
	eServ "github.com/mhd53/quanta-fitness-server/internal/exercise"
	wServ "github.com/mhd53/quanta-fitness-server/internal/workout"
	"github.com/mhd53/quanta-fitness-server/pkg/format"
)

var (
	service    serv.EsetService
	notImplErr = errors.New("Not Implemented!")
	// deniedErr = errors.New("Access Denied!")
	internalErr = errors.New("Internal Error!")
)

type server struct{}

type EsetServer interface {
	AddEsetToExercise(uname, eid string, actualRC int, dur, restDur float32) (entity.Eset, error)
	UpdateEset(id, uname string, updates entity.EsetUpdate) (bool, error)
	DeleteEset(id, uname string) (bool, error)
	GetEset(id, uname string) (entity.Eset, error)
	GetEsetsForExercise(eid, uname string) ([]entity.Eset, error)
}

func NewEsetServer(us ustore.UserStore, ess esstore.EsetStore, es estore.ExerciseStore, ws wstore.WorkoutStore) EsetServer {
	wAuthorizer := wServ.NewWorkoutAuthorizer(ws, us)
	eAuthorizer := eServ.NewExerciseAuthorizer(es, us, wAuthorizer)
	authorizer := serv.NewEsetAuthorizer(ess, us, eAuthorizer)
	validator := serv.NewEsetValidator()
	service = serv.NewEsetService(ess, authorizer, validator)
	return &server{}
}

func (*server) AddEsetToExercise(uname, eid string, actualRC int, dur, restDur float32) (entity.Eset, error) {
	intID, err := format.ConvertToBase64(eid)
	if err != nil {
		log.Panic("API Error: ", err.Error())
		return entity.Eset{}, internalErr
	}

	return service.AddEsetToExercise(uname, intID, actualRC, dur, restDur)
}

func (*server) UpdateEset(id, uname string, updates entity.EsetUpdate) (bool, error) {
	intID, err := format.ConvertToBase64(id)
	if err != nil {
		log.Panic("API Error: ", err.Error())
		return false, internalErr
	}

	err2 := service.UpdateEset(intID, uname, updates)
	if err2 != nil {
		log.Panic("API Error: ", err2.Error())
		return false, internalErr
	}

	return true, nil

}

func (*server) DeleteEset(id, uname string) (bool, error) {
	intID, err := format.ConvertToBase64(id)
	if err != nil {
		log.Panic("API Error: ", err.Error())
		return false, internalErr
	}

	err2 := service.DeleteEset(intID, uname)
	if err2 != nil {
		log.Panic("API Error: ", err2.Error())
		return false, internalErr
	}

	return true, nil

}

func (*server) GetEset(id, uname string) (entity.Eset, error) {
	intID, err := format.ConvertToBase64(id)
	if err != nil {
		log.Panic("API Error: ", err.Error())
		return entity.Eset{}, internalErr
	}

	got, err2 := service.GetEset(intID, uname)
	if err2 != nil {
		log.Panic("API Error: ", err2.Error())
		return entity.Eset{}, internalErr
	}
	return got, nil

}

func (*server) GetEsetsForExercise(eid, uname string) ([]entity.Eset, error) {
	var esets []entity.Eset

	intID, err := format.ConvertToBase64(eid)
	if err != nil {
		log.Panic("API Error: ", err.Error())
		return esets, internalErr
	}

	esets, err2 := service.GetEsetsForExercise(intID, uname)
	if err2 != nil {
		log.Panic("API Error: ", err2.Error())
		return esets, internalErr
	}

	return esets, nil
}
