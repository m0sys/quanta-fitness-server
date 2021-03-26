package exercise

type ExerciseAuth interface {
	AuthorizeReadAccess(uname string) (bool, error)
	AuthorizeWorkoutAccess(uname string, wid int64) (bool, error)
	AuthorizeExerciseAccess(uname string, eid int64) (bool, error)
}
