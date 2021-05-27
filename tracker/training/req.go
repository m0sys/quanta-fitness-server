package training

type FetchWorkoutLogExerciseLogsReq struct {
	AthleteID    string
	WorkoutLogID string
}

func ValidateFetchWorkoutLogExerciseLogsReq(req FetchWorkoutLogExerciseLogsReq) error {
	if req.AthleteID == "" {
		return ErrAthleteIDCannotBeEmpty
	}

	if req.WorkoutLogID == "" {
		return ErrWorkoutLogIDCannotBeEmpty
	}

	return nil
}

type AddSetLogToExerciseLogReq struct {
	AthleteID      string
	WorkoutLogID   string
	ExerciseLogID  string
	ActualRepCount int
	Duration       float64
}

func ValidateAddSetLogToExerciseLogReq(req AddSetLogToExerciseLogReq) error {
	if req.AthleteID == "" {
		return ErrAthleteIDCannotBeEmpty
	}

	if req.WorkoutLogID == "" {
		return ErrWorkoutLogIDCannotBeEmpty
	}

	if req.ExerciseLogID == "" {
		return ErrExerciseLogIDCannotBeEmpty
	}

	if req.ActualRepCount == 0 {
		return ErrActualRepCountCannotBeEmpty
	}

	if req.Duration == 0 {
		return ErrDurationCannotBeEmpty
	}

	return nil
}

type FetchSetLogsForExerciseLogReq struct {
	AthleteID     string
	WorkoutLogID  string
	ExerciseLogID string
}

func ValidateFetchSetLogsForExerciseLogReq(req FetchSetLogsForExerciseLogReq) error {
	if req.AthleteID == "" {
		return ErrAthleteIDCannotBeEmpty
	}

	if req.WorkoutLogID == "" {
		return ErrWorkoutLogIDCannotBeEmpty
	}

	if req.ExerciseLogID == "" {
		return ErrExerciseLogIDCannotBeEmpty
	}

	return nil
}

type MoveToNextExerciseLogReq struct {
	AthleteID    string
	WorkoutLogID string
}

func ValidateMoveToNextExerciseLogReq(req MoveToNextExerciseLogReq) error {
	if req.AthleteID == "" {
		return ErrAthleteIDCannotBeEmpty
	}

	if req.WorkoutLogID == "" {
		return ErrWorkoutLogIDCannotBeEmpty
	}

	return nil
}

type FetchCurrentExerciseLogReq struct {
	AthleteID    string
	WorkoutLogID string
}

func ValidateFetchCurrentExerciseLogReq(req FetchCurrentExerciseLogReq) error {
	if req.AthleteID == "" {
		return ErrAthleteIDCannotBeEmpty
	}

	if req.WorkoutLogID == "" {
		return ErrWorkoutLogIDCannotBeEmpty
	}

	return nil
}

type RemoveWorkoutLogReq struct {
	AthleteID    string
	WorkoutLogID string
}

func ValidateRemoveWorkoutLogReq(req RemoveWorkoutLogReq) error {
	if req.AthleteID == "" {
		return ErrAthleteIDCannotBeEmpty
	}

	if req.WorkoutLogID == "" {
		return ErrWorkoutLogIDCannotBeEmpty
	}

	return nil
}
