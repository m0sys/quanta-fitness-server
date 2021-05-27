package training

import (
	"log"

	el "github.com/mhd53/quanta-fitness-server/tracker/exerciselog"
	sl "github.com/mhd53/quanta-fitness-server/tracker/setlog"
	wl "github.com/mhd53/quanta-fitness-server/tracker/workoutlog"
)

type TrainingService struct {
	repo Repository
}

func NewTrainingService(repository Repository) TrainingService {
	return TrainingService{repo: repository}
}

func (t TrainingService) FetchWorkoutLogs(aid string) ([]WorkoutLogRes, error) {
	var results []WorkoutLogRes

	wlogs, err := t.repo.FindAllWorkoutLogsForAthlete(aid)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return results, errInternal
	}

	for _, wlog := range wlogs {
		res := mapWorkoutLogToWorkoutLogRes(wlog)
		results = append(results, res)
	}

	return results, nil
}

func (t TrainingService) FetchWorkoutLogExerciseLogs(
	req FetchWorkoutLogExerciseLogsReq,
) ([]ExerciseLogRes, error) {
	var results []ExerciseLogRes

	if err := t.validateWlog(req.AthleteID, req.WorkoutLogID); err != nil {
		return results, err
	}

	wlog, err := t.findWorkoutLog(req.WorkoutLogID)
	if err != nil {
		return results, err
	}

	elogs, err := t.repo.FindAllExerciseLogsForWorkoutLog(wlog)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return results, errInternal
	}

	for _, elog := range elogs {
		res := mapElogToElogRes(elog)
		results = append(results, res)
	}

	return results, nil
}

func (t TrainingService) validateWlog(aid, wlid string) error {
	wlog, err := t.findWorkoutLog(wlid)
	if err != nil {
		return err
	}

	if !isAuthorizedWL(aid, wlog) {
		return ErrUnauthorizedAccess
	}

	return nil
}

func (t TrainingService) findWorkoutLog(id string) (wl.WorkoutLog, error) {
	wlog, found, err := t.repo.FindWorkoutLogByID(id)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return wl.WorkoutLog{}, errInternal
	}

	if !found {
		return wl.WorkoutLog{}, ErrWorkoutLogNotFound
	}

	return wlog, nil
}

func isAuthorizedWL(aid string, wlog wl.WorkoutLog) bool {
	return wlog.AthleteID() == aid
}

func (t TrainingService) AddSetLogToExerciseLog(
	req AddSetLogToExerciseLogReq,
) error {
	if err := t.validateWlog(req.AthleteID, req.WorkoutLogID); err != nil {
		return err
	}

	if err := t.validateElog(req.AthleteID, req.WorkoutLogID, req.ExerciseLogID); err != nil {
		return err
	}

	elog, err := t.findExerciseLog(req.ExerciseLogID)
	if err != nil {
		return err
	}

	slogs, err := t.repo.FindAllSetLogsForExerciseLog(elog)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	eMetrics := elog.Metrics()
	if len(slogs) == eMetrics.NumSets() {
		return ErrCannotExceedNumSets
	}

	metrics := sl.NewMetrics(req.ActualRepCount, req.Duration)
	slog := sl.NewSetLog(elog.ID(), metrics)
	err = t.repo.StoreSetLog(slog)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	return nil
}

func (t TrainingService) validateElog(aid, wlid, elid string) error {
	elog, err := t.findExerciseLog(elid)
	if err != nil {
		return err
	}

	wlog, err := t.findWorkoutLog(wlid)
	if err != nil {
		return err
	}

	if !isAuthorizedEL(aid, wlog, elog) {
		return ErrUnauthorizedAccess
	}

	return nil
}

func (t TrainingService) findExerciseLog(id string) (el.ExerciseLog, error) {
	elog, found, err := t.repo.FindExerciseLogByID(id)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return el.ExerciseLog{}, errInternal
	}

	if !found {
		return el.ExerciseLog{}, ErrExerciseLogNotFound
	}

	return elog, nil
}

func isAuthorizedEL(aid string, wlog wl.WorkoutLog, elog el.ExerciseLog) bool {
	return aid == wlog.AthleteID() && elog.WorkoutLogID() == wlog.ID()
}

func (t TrainingService) MoveToNextExerciseLog(
	req MoveToNextExerciseLogReq,
) error {
	if err := ValidateMoveToNextExerciseLogReq(req); err != nil {
		return err
	}

	if err := t.validateWlog(req.AthleteID, req.WorkoutLogID); err != nil {
		return err
	}

	wlog, err := t.findWorkoutLog(req.WorkoutLogID)
	if err != nil {
		return err
	}

	if wlog.Completed() {
		return ErrWorkoutLogAlreadyCompleted
	}

	elogs, err := t.repo.FindAllExerciseLogsForWorkoutLog(wlog)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	wlog.NextPos()
	if wlog.CurrentPos() == len(elogs) {
		wlog.Complete()
	}

	err = t.repo.UpdateWorkoutLog(wlog)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	return nil
}

func (t TrainingService) FetchCurrentExerciseLog(req FetchCurrentExerciseLogReq) (ExerciseLogRes, error) {
	if err := ValidateFetchCurrentExerciseLogReq(req); err != nil {
		return ExerciseLogRes{}, err
	}

	if err := t.validateWlog(req.AthleteID, req.WorkoutLogID); err != nil {
		return ExerciseLogRes{}, err
	}

	wlog, err := t.findWorkoutLog(req.WorkoutLogID)
	if err != nil {
		return ExerciseLogRes{}, err
	}

	elogs, err := t.repo.FindAllExerciseLogsForWorkoutLog(wlog)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return ExerciseLogRes{}, errInternal
	}
	res := mapElogToElogRes(elogs[wlog.CurrentPos()])

	return res, nil
}

func (t TrainingService) FetchSetLogsForExerciseLog(req FetchSetLogsForExerciseLogReq) ([]SetLogRes, error) {
	var results []SetLogRes

	if err := ValidateFetchSetLogsForExerciseLogReq(req); err != nil {
		return results, err
	}

	if err := t.validateWlog(req.AthleteID, req.WorkoutLogID); err != nil {
		return results, err
	}

	if err := t.validateElog(req.AthleteID, req.WorkoutLogID, req.ExerciseLogID); err != nil {
		return results, err
	}

	elog, err := t.findExerciseLog(req.ExerciseLogID)
	if err != nil {
		return results, err
	}

	slogs, err := t.repo.FindAllSetLogsForExerciseLog(elog)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return results, errInternal
	}

	for _, slog := range slogs {
		res := mapSetLogToSetLogRes(slog)
		results = append(results, res)
	}

	return results, nil
}

func (t TrainingService) RemoveWorkoutLog(req RemoveWorkoutLogReq) error {
	if err := ValidateRemoveWorkoutLogReq(req); err != nil {
		return err
	}
	if err := t.validateWlog(req.AthleteID, req.WorkoutLogID); err != nil {
		return err
	}

	wlog, err := t.findWorkoutLog(req.WorkoutLogID)
	if err != nil {
		return err
	}

	elogs, err := t.repo.FindAllExerciseLogsForWorkoutLog(wlog)
	if err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	for _, elog := range elogs {
		slogs, err := t.repo.FindAllSetLogsForExerciseLog(elog)
		if err != nil {
			log.Printf("%s: %s", errSlur, err.Error())
			return errInternal
		}

		for _, slog := range slogs {
			if err := t.repo.RemoveSetLog(slog); err != nil {
				log.Printf("%s: %s", errSlur, err.Error())
				return errInternal
			}
		}

		if err := t.repo.RemoveExerciseLog(elog); err != nil {
			log.Printf("%s: %s", errSlur, err.Error())
			return errInternal
		}
	}

	if err := t.repo.RemoveWorkoutLog(wlog); err != nil {
		log.Printf("%s: %s", errSlur, err.Error())
		return errInternal
	}

	return nil
}

// Map funcs.

func mapWorkoutLogToWorkoutLogRes(wlog wl.WorkoutLog) WorkoutLogRes {
	return WorkoutLogRes{
		ID:         wlog.ID(),
		Title:      wlog.Title(),
		Date:       wlog.Date(),
		CurrentPos: wlog.CurrentPos(),
		Completed:  wlog.Completed(),
	}
}

func mapElogToElogRes(elog el.ExerciseLog) ExerciseLogRes {
	metrics := elog.Metrics()
	return ExerciseLogRes{
		ID:           elog.ID(),
		WorkoutLogID: elog.WorkoutLogID(),
		Name:         elog.Name(),
		TargetRep:    metrics.TargetRep(),
		NumSets:      metrics.NumSets(),
		Weight:       metrics.Weight(),
		RestDur:      metrics.RestDur(),
		Completed:    elog.Completed(),
		Pos:          elog.Pos(),
	}
}

func mapSetLogToSetLogRes(slog sl.SetLog) SetLogRes {
	metrics := slog.Metrics()

	return SetLogRes{
		ID:             slog.ID(),
		ExerciseLogID:  slog.ExerciseLogID(),
		ActualRepCount: metrics.ActualRepCount(),
		Duration:       metrics.Dur(),
	}
}
