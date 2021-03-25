package entity

import (
	"time"
)

type BaseEset struct {
	Username          int64   `json:"username"`
	EID               int64   `json:"eid"`
	ActualRepCount    int     `json:"actual_rep_count"`
	Duraction         float32 `json:"duraction"`
	RestTimeDuraction float32 `json:"rest_time_dur"`
}

type Eset struct {
	BaseEset
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
