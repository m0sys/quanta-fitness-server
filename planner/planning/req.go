package planning

type CreateNewWorkoutPlanReq struct {
	AthleteID string
	Title     string
}

func ValidateCreateNewWorkoutPlanReq(req CreateNewWorkoutPlanReq) error {
	if req.AthleteID == "" {
		return ErrAthleteIDCannotBeEmpty
	}

	if req.Title == "" {
		return ErrTitleCannotBeEmpty
	}

	return nil
}

type AddNewExerciseToWorkoutPlanReq struct {
	AthleteID     string
	WorkoutPlanID string
	Name          string
	TargetRep     int
	NumSets       int
	Weight        float64
	RestDur       float64
}

func ValidateAddExerciseToWorkoutPlanReq(req AddNewExerciseToWorkoutPlanReq) error {
	if req.AthleteID == "" {
		return ErrAthleteIDCannotBeEmpty
	}

	if req.WorkoutPlanID == "" {
		return ErrWorkoutPlanIDCannotBeEmpty
	}

	if req.Name == "" {
		return ErrNameCannotBeEmpty
	}

	if req.TargetRep == 0 {
		return ErrTargetRepCannotBeEmpty
	}

	if req.NumSets == 0 {
		return ErrNumSetsCannotBeEmpty
	}

	return nil
}

type RemoveExerciseFromWorkoutPlanReq struct {
	AthleteID     string
	WorkoutPlanID string
	ExerciseID    string
}

func ValidateRemoveExerciseFromWorkoutPlanReq(req RemoveExerciseFromWorkoutPlanReq) error {
	if req.AthleteID == "" {
		return ErrAthleteIDCannotBeEmpty
	}

	if req.WorkoutPlanID == "" {
		return ErrWorkoutPlanIDCannotBeEmpty
	}

	if req.ExerciseID == "" {
		return ErrExerciseIDCannotBeEmpty
	}

	return nil
}

type FetchWorkoutPlanExercisesReq struct {
	AthleteID     string
	WorkoutPlanID string
}

func ValidateFetchWorkoutPlanExercisesReq(req FetchWorkoutPlanExercisesReq) error {
	if req.AthleteID == "" {
		return ErrAthleteIDCannotBeEmpty
	}

	if req.WorkoutPlanID == "" {
		return ErrWorkoutPlanIDCannotBeEmpty
	}

	return nil
}

type EditWorkoutPlanTitleReq struct {
	AthleteID     string
	WorkoutPlanID string
	Title         string
}

func ValidateEditWorkoutPlanTitleReq(req EditWorkoutPlanTitleReq) error {
	if req.AthleteID == "" {
		return ErrAthleteIDCannotBeEmpty
	}

	if req.WorkoutPlanID == "" {
		return ErrWorkoutPlanIDCannotBeEmpty
	}

	if req.Title == "" {
		return ErrTitleCannotBeEmpty
	}

	return nil
}

type EditExerciseNameReq struct {
	AthleteID     string
	WorkoutPlanID string
	ExerciseID    string
	Name          string
}

func ValidateEditExerciseNameReq(req EditExerciseNameReq) error {
	if req.AthleteID == "" {
		return ErrAthleteIDCannotBeEmpty
	}

	if req.WorkoutPlanID == "" {
		return ErrWorkoutPlanIDCannotBeEmpty
	}

	if req.ExerciseID == "" {
		return ErrExerciseIDCannotBeEmpty
	}

	if req.Name == "" {
		return ErrNameCannotBeEmpty
	}

	return nil
}

type EditExerciseTargetRepReq struct {
	AthleteID     string
	WorkoutPlanID string
	ExerciseID    string
	TargetRep     int
}

func ValidateEditExerciseTargetRepReq(req EditExerciseTargetRepReq) error {
	if req.AthleteID == "" {
		return ErrAthleteIDCannotBeEmpty
	}

	if req.WorkoutPlanID == "" {
		return ErrWorkoutPlanIDCannotBeEmpty
	}

	if req.ExerciseID == "" {
		return ErrExerciseIDCannotBeEmpty
	}

	if req.TargetRep == 0 {
		return ErrTargetRepCannotBeEmpty
	}

	return nil
}

type EditExerciseNumSetsReq struct {
	AthleteID     string
	WorkoutPlanID string
	ExerciseID    string
	NumSets       int
}

func ValidateEditExerciseNumSetsReq(req EditExerciseNumSetsReq) error {
	if req.AthleteID == "" {
		return ErrAthleteIDCannotBeEmpty
	}

	if req.WorkoutPlanID == "" {
		return ErrWorkoutPlanIDCannotBeEmpty
	}

	if req.ExerciseID == "" {
		return ErrExerciseIDCannotBeEmpty
	}

	if req.NumSets == 0 {
		return ErrNumSetsCannotBeEmpty
	}

	return nil
}

type EditExerciseWeightReq struct {
	AthleteID     string
	WorkoutPlanID string
	ExerciseID    string
	Weight        float64
}

func ValidateEditExerciseWeightReq(req EditExerciseWeightReq) error {
	if req.AthleteID == "" {
		return ErrAthleteIDCannotBeEmpty
	}

	if req.WorkoutPlanID == "" {
		return ErrWorkoutPlanIDCannotBeEmpty
	}

	if req.ExerciseID == "" {
		return ErrExerciseIDCannotBeEmpty
	}

	return nil
}

type EditExerciseRestDurReq struct {
	AthleteID     string
	WorkoutPlanID string
	ExerciseID    string
	RestDur       float64
}

func ValidateEditExerciseRestDurReq(req EditExerciseRestDurReq) error {
	if req.AthleteID == "" {
		return ErrAthleteIDCannotBeEmpty
	}

	if req.WorkoutPlanID == "" {
		return ErrWorkoutPlanIDCannotBeEmpty
	}

	if req.ExerciseID == "" {
		return ErrExerciseIDCannotBeEmpty
	}

	return nil
}

type RemoveWorkoutPlanReq struct {
	AthleteID     string
	WorkoutPlanID string
}

func ValidateRemoveWorkoutPlanReq(req RemoveWorkoutPlanReq) error {
	if req.AthleteID == "" {
		return ErrAthleteIDCannotBeEmpty
	}

	if req.WorkoutPlanID == "" {
		return ErrWorkoutPlanIDCannotBeEmpty
	}

	return nil
}
