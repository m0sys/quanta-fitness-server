// Entity contains the Exercise entity.
package exercise

import (
	"time"

	"github.com/mhd53/quanta-fitness-server/internal/id"
)

type Exercise struct {
	ID        id.ID
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
