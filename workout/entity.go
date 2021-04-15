// Entity contains the Workout entity.
package workout

import "time"

type Workout struct {
	ID        int64
	AID       int64
	Title     string
	Date      time.Time
	Exercises []Exercise
	State     string // 'incomplete', 'inprogress', 'done'
	Duration  time.Duration
}
