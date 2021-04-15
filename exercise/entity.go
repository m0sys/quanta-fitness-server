// Entity contains the Exercise entity.
package exercise

import "time"

type Exercise struct {
	ID        int64
	WID       int64
	Name      string
	RestTime  float64
	Distance  float64
	Weight    float64
	Duration  time.Duration
	TargetRep int
	Sets      []Set
	State     string // 'incomplete', 'inprogress', 'done'
}
