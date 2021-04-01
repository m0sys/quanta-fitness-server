package workouttests

import (
	"time"

	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

const (
	MOCK_UNAME       = "robin"
	MOCK_TITLE       = "Chest Day"
	MOCK_VAL_TITLE   = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus."
	MOCK_INVAL_TITLE = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vivamus.."
)

func CreateValidMockWorkout(id int64) entity.Workout {

	workout := entity.Workout{
		BaseWorkout: entity.BaseWorkout{
			Title:    MOCK_TITLE,
			Username: MOCK_UNAME,
		},
		ID:        string(id),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return workout
}

func CreateValidMockBaseWorkout() entity.BaseWorkout {
	return entity.BaseWorkout{
		Title:    MOCK_VAL_TITLE,
		Username: MOCK_UNAME,
	}
}

func CreateInvalidMockBaseWorkout() entity.BaseWorkout {
	return entity.BaseWorkout{
		Title:    MOCK_INVAL_TITLE,
		Username: MOCK_UNAME,
	}
}
