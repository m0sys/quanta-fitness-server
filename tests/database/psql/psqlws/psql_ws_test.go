package psqlwstest

import (
	// "database/sql"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/mhd53/quanta-fitness-server/database/psql/psqlus"
	"github.com/mhd53/quanta-fitness-server/database/psql/psqlws"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
	// us "github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
	// ws "github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"
)

// func setup() (us.UserStore, ws.WorkoutStore) {
// 	psqlUS := psqlus.NewPsqlUserStore()
// 	psqlUS := psqlws.NewPsqlWorkoutStore()
//
// 	user := entity.BaseUser{
// 		Username: "robin",
// 		Email:    "robin@gmail.com",
// 		Password: "robinhood",
// 	}
// 	newUser, err := psqlDB.Save(user)
// }
//
// func clean() {
// }

func TestSave(t *testing.T) {
	psqlUS := psqlus.NewPsqlUserStore()
	user := entity.BaseUser{
		Username: "robin",
		Email:    "robin@gmail.com",
		Password: "robinhood",
	}
	newUser, err := psqlUS.Save(user)

	assert.Nil(t, err)
	assert.NotEmpty(t, newUser)
	assert.Equal(t, "robin", newUser.Username)

	psqlWS := psqlws.NewPsqlWorkoutStore()

	workout := entity.BaseWorkout{
		Username: "robin",
		Title:    "Chest Day!",
	}

	newWorkout, err := psqlWS.Save(workout)

	assert.Nil(t, err)
	assert.NotEmpty(t, newWorkout)
	assert.Equal(t, "robin", newWorkout.Username)
	assert.Equal(t, "Chest Day!", newWorkout.Title)

	t.Cleanup(func() {
		psqlWS.DeleteWorkout(newWorkout.ID)
		psqlUS.DeleteUser(newUser.ID)
	})
}
