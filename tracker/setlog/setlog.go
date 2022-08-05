package setlog

import (
	"github.com/m0sys/quanta-fitness-server/pkg/uuid"
	"github.com/m0sys/quanta-fitness-server/units"
)

type SetLog struct {
	uuid    string
	elid    string
	metrics Metrics
}

type Metrics struct {
	actualRepCount int
	duration       units.Second
}

// This constructor should only be used through SetPlanToSetLogTranslator.
func NewSetLog(elid string, metrics Metrics) SetLog {
	return SetLog{
		uuid:    uuid.GenerateUUID(),
		elid:    elid,
		metrics: metrics,
	}
}

// This should only be used to restore data from persistence layer.
func RestoreSetLog(id, elid string, metrics Metrics) SetLog {
	return SetLog{
		uuid:    id,
		elid:    elid,
		metrics: metrics,
	}
}

func (s *SetLog) ID() string {
	return s.uuid
}

func (s *SetLog) ExerciseLogID() string {
	return s.elid
}

func (s *SetLog) Metrics() Metrics {
	return s.metrics
}

func NewMetrics(actualRepCount int, dur float64) Metrics {
	return Metrics{
		actualRepCount: actualRepCount,
		duration:       units.Second(dur),
	}
}

func (m *Metrics) ActualRepCount() int {
	return m.actualRepCount
}

func (m *Metrics) Dur() units.Second {
	return m.duration
}
