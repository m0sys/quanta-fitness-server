package psqlwstest

import (
	// "database/sql"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/mhd53/quanta-fitness-server/database/psql/psqlus"
	"github.com/mhd53/quanta-fitness-server/database/psql/psqlws"
	us "github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
	ws "github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

func TestSave(t *testing.T) {
	psqlUS, psqlWS, uid := setup(t)

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
		psqlUS.DeleteUser(uid)
	})
}

func TestUpdate(t *testing.T) {
	psqlUS, psqlWS, uid := setup(t)

	workout := entity.BaseWorkout{
		Username: "robin",
		Title:    "Chest Day!",
	}

	newWorkout, err := psqlWS.Save(workout)

	assert.Nil(t, err)
	assert.NotEmpty(t, newWorkout)
	assert.Equal(t, "robin", newWorkout.Username)
	assert.Equal(t, "Chest Day!", newWorkout.Title)

	updates := entity.BaseWorkout{
		Username: "robin",
		Title:    "Chest Day 2!",
	}
	err = psqlWS.Update(newWorkout.ID, updates)
	assert.Nil(t, err)

	t.Cleanup(func() {
		psqlWS.DeleteWorkout(newWorkout.ID)
		psqlUS.DeleteUser(uid)
	})
}

func TestFindWorkoutById(t *testing.T) {
	psqlUS, psqlWS, uid := setup(t)

	workout := entity.BaseWorkout{
		Username: "robin",
		Title:    "Chest Day!",
	}

	newWorkout, err := psqlWS.Save(workout)

	assert.Nil(t, err)
	assert.NotEmpty(t, newWorkout)
	assert.Equal(t, "robin", newWorkout.Username)
	assert.Equal(t, "Chest Day!", newWorkout.Title)

	got, found, err := psqlWS.FindWorkoutById(newWorkout.ID)

	assert.Nil(t, err)
	assert.True(t, found)
	assert.Equal(t, "robin", got.Username)
	assert.Equal(t, "Chest Day!", got.Title)
	assert.Equal(t, newWorkout.ID, got.ID)

	t.Cleanup(func() {
		psqlWS.DeleteWorkout(newWorkout.ID)
		psqlUS.DeleteUser(uid)
	})
}

func TestDelete(t *testing.T) {
	psqlUS, psqlWS, uid := setup(t)

	workout := entity.BaseWorkout{
		Username: "robin",
		Title:    "Chest Day!",
	}

	newWorkout, err := psqlWS.Save(workout)

	assert.Nil(t, err)
	assert.NotEmpty(t, newWorkout)
	assert.Equal(t, "robin", newWorkout.Username)
	assert.Equal(t, "Chest Day!", newWorkout.Title)

	err = psqlWS.DeleteWorkout(newWorkout.ID)
	assert.Nil(t, err)

	t.Cleanup(func() {
		psqlUS.DeleteUser(uid)
	})
}

func TestFindAllWorkoutsByUname(t *testing.T) {
	psqlUS, psqlWS, uid := setup(t)

	workout := entity.BaseWorkout{
		Username: "robin",
		Title:    "Chest Day!",
	}

	count := 3
	for i := 0; i < count; i++ {
		newWorkout, err := psqlWS.Save(workout)
		assert.Nil(t, err)
		assert.NotEmpty(t, newWorkout)
		assert.Equal(t, "robin", newWorkout.Username)
		assert.Equal(t, "Chest Day!", newWorkout.Title)

	}

	got, err := psqlWS.FindAllWorkoutsByUname("robin")
	assert.Nil(t, err)
	assert.NotEmpty(t, got)
	assert.Equal(t, count, len(got))
	assert.Equal(t, "Chest Day!", got[0].BaseWorkout.Title)
	assert.Equal(t, "robin", got[0].BaseWorkout.Username)

	t.Cleanup(func() {
		psqlUS.DeleteUser(uid)
		for i := 0; i < count; i++ {
			err := psqlWS.DeleteWorkout(got[i].ID)
			assert.Nil(t, err)
		}
	})
}

// Util funcs.

func setup(t *testing.T) (us.UserStore, ws.WorkoutStore, int64) {
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
	return psqlUS, psqlWS, newUser.ID
}
