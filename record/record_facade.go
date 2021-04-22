// RecordFacade provides a facade to access all use cases for storing workout information.
package record

import (
	"github.com/mhd53/quanta-fitness-server/account"
)

type RecordFacade interface {
	workoutInteractor
}

var (
	wGateway WorkoutGateway
	aGateway AthleteGateway
	uGateway account.AccountGateway
)

type interactor struct{}

func NewRecordFacade(workoutGateway WorkoutGateway, athleteGateway, accountGateway account.AccountGateway) RecordInteractor {
	wGateway = workoutGateway
	uGateway = accountGateway
	aGateway = athleteGateway
	return &interactor{}
}
