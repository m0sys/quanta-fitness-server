package manager

import "time"

// WeightRecordRes ds for FetchWeightHistory response.
type WeightRecordRes struct {
	Amount float64 // in kg
	Date   time.Time
}
