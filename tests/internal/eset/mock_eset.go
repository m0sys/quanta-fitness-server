package esettest

import (
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

func CreateValidBaseRobinSet() entity.BaseEset {
	return entity.BaseEset{
		Username: "robin",
		EID:      0,
		SMetric: entity.SMetric{
			ActualRepCount:    6,
			Duraction:         183.0,
			RestTimeDuraction: 120.2,
		},
	}
}

func CreateValidBaseBobinSet() entity.BaseEset {
	return entity.BaseEset{
		Username: "bobin",
		EID:      0,
		SMetric: entity.SMetric{
			ActualRepCount:    6,
			Duraction:         183.0,
			RestTimeDuraction: 120.2,
		},
	}
}
