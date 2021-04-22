// Entity contains the Workout entity.
package workout

import (
	"time"

	"github.com/mhd53/quanta-fitness-server/internal/id"
)

type Workout struct {
	ID       id.ID
	AID      id.ID
	Title    string
	Date     time.Time
	State    string // 'incomplete', 'inprogress', 'done'
	Duration time.Duration
}
