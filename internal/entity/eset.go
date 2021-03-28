package entity

import (
	"time"
)

// The only time when you should be creating a exercise set is when you are
// executing the workout. That's why SMetric is required @BaseEset.

type BaseEset struct {
	Username string `json:"username"`
	EID      int64  `json:"eid"`
	SMetric
}

type SMetric struct {
	ActualRepCount   int     `json:"actual_rep_count"`
	Duration         float32 `json:"duration"`
	RestTimeDuration float32 `json:"rest_time_dur"`
}

type EsetUpdate struct {
	SMetric
}

type Eset struct {
	BaseEset
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
