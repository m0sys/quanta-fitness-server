package exercisetest

import (
	"github.com/mhd53/quanta-fitness-server/internal/entity"
	ats "github.com/mhd53/quanta-fitness-server/tests/internal/auth"
)

const (
	MOCK_VALID_NAME   = "Lorem ipsum dolor sit amet porta ante."
	MOCK_INVALID_NAME = "Lorem ipsum dolor sit amet porta ante.."
)

func CreateMockValidBaseExercise() entity.BaseExercise {
	return entity.BaseExercise{
		Name:     MOCK_VALID_NAME,
		WID:      0,
		Username: ats.MOCK_USERNAME,
	}
}

func CreateMockInvalidBaseExercise() entity.BaseExercise {
	return entity.BaseExercise{
		Name:     MOCK_INVALID_NAME,
		WID:      0,
		Username: ats.MOCK_USERNAME,
	}
}

func CreateMockValidUpdateExercise() entity.ExerciseUpdate {
	return entity.ExerciseUpdate{
		Name: MOCK_VALID_NAME,
		Metrics: entity.Metrics{
			Weight:    192.0,
			TargetRep: 5,
			RestTime:  120.0,
			NumSets:   3,
		},
	}
}

func CreateMockInvalidUpdateExercise() entity.ExerciseUpdate {
	return entity.ExerciseUpdate{
		Name: MOCK_INVALID_NAME,
		Metrics: entity.Metrics{
			Weight:    192.0,
			TargetRep: 5,
			RestTime:  120.0,
			NumSets:   3,
		},
	}
}
