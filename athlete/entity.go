// Entity contains the Athlete entity.
package athlete

import "time"

type Athlete struct {
	AID           int64
	Workouts      map[time.Time]Workout
	Height        float64
	CurrentWeight Weight
	WeightHistory []Weight
}

type Weight struct {
	Amount float64
	Date   time.Time
}
