package esettest

import (
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

func CreateValidBaseRobinSet() entity.BaseEset {
	return entity.BaseEset{
		Username: "robin",
		EID:      "0",
		SMetric: entity.SMetric{
			ActualRepCount:   6,
			Duration:         183.0,
			RestTimeDuration: 120.2,
		},
	}
}

func CreateValidBaseBobinSet() entity.BaseEset {
	return entity.BaseEset{
		Username: "bobin",
		EID:      "0",
		SMetric: entity.SMetric{
			ActualRepCount:   6,
			Duration:         183.0,
			RestTimeDuration: 120.2,
		},
	}
}

func CreateInvalidEsetUpdate() entity.EsetUpdate {
	return entity.EsetUpdate{
		SMetric: entity.SMetric{

			ActualRepCount:   -5,
			Duration:         120.0,
			RestTimeDuration: 123.3,
		},
	}

}

func CreateValidEsetUpdate() entity.EsetUpdate {
	return entity.EsetUpdate{
		SMetric: entity.SMetric{

			ActualRepCount:   10,
			Duration:         220.0,
			RestTimeDuration: 323.3,
		},
	}

}
