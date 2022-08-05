package psqlwstest

import (
	// "database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/m0sys/quanta-fitness-server/database/psql/psqlus"
	"github.com/m0sys/quanta-fitness-server/database/psql/psqlws"
	us "github.com/m0sys/quanta-fitness-server/internal/datastore/userstore"
	ws "github.com/m0sys/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/m0sys/quanta-fitness-server/internal/entity"
)

func TestSave(t *testing.T) {
	psqlUS, psqlWS, uid := setup(t)

	workout := entity.BaseWorkout{
		Username: "bobin",
		Title:    "Chest Day!",
	}

	newWorkout, err := psqlWS.Save(workout)

	assert.Nil(t, err)
	assert.NotEmpty(t, newWorkout)
	assert.Equal(t, "bobin", newWorkout.Username)
	assert.Equal(t, "Chest Day!", newWorkout.Title)

	t.Cleanup(func() {
		err = psqlWS.DeleteWorkout(newWorkout.ID)
		assert.Nil(t, err)
		_, err = psqlUS.DeleteUser(uid)
		assert.Nil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	psqlUS, psqlWS, uid := setup(t)

	workout := entity.BaseWorkout{
		Username: "bobin",
		Title:    "Chest Day!",
	}

	newWorkout, err := psqlWS.Save(workout)

	assert.Nil(t, err)
	assert.NotEmpty(t, newWorkout)
	assert.Equal(t, "bobin", newWorkout.Username)
	assert.Equal(t, "Chest Day!", newWorkout.Title)

	updates := entity.BaseWorkout{
		Username: "bobin",
		Title:    "Chest Day 2!",
	}
	err = psqlWS.Update(newWorkout.ID, updates)
	assert.Nil(t, err)

	t.Cleanup(func() {
		err = psqlWS.DeleteWorkout(newWorkout.ID)
		assert.Nil(t, err)
		_, err = psqlUS.DeleteUser(uid)
		assert.Nil(t, err)
	})
}

func TestFindWorkoutById(t *testing.T) {
	psqlUS, psqlWS, uid := setup(t)

	workout := entity.BaseWorkout{
		Username: "bobin",
		Title:    "Chest Day!",
	}

	newWorkout, err := psqlWS.Save(workout)

	assert.Nil(t, err)
	assert.NotEmpty(t, newWorkout)
	assert.Equal(t, "bobin", newWorkout.Username)
	assert.Equal(t, "Chest Day!", newWorkout.Title)

	got, found, err := psqlWS.FindWorkoutById(newWorkout.ID)

	assert.Nil(t, err)
	assert.True(t, found)
	assert.Equal(t, "bobin", got.Username)
	assert.Equal(t, "Chest Day!", got.Title)
	assert.Equal(t, newWorkout.ID, got.ID)

	t.Cleanup(func() {
		err = psqlWS.DeleteWorkout(newWorkout.ID)
		assert.Nil(t, err)
		_, err = psqlUS.DeleteUser(uid)
		assert.Nil(t, err)
	})
}

func TestDelete(t *testing.T) {
	psqlUS, psqlWS, uid := setup(t)

	workout := entity.BaseWorkout{
		Username: "bobin",
		Title:    "Chest Day!",
	}

	newWorkout, err := psqlWS.Save(workout)

	assert.Nil(t, err)
	assert.NotEmpty(t, newWorkout)
	assert.Equal(t, "bobin", newWorkout.Username)
	assert.Equal(t, "Chest Day!", newWorkout.Title)

	err = psqlWS.DeleteWorkout(newWorkout.ID)
	assert.Nil(t, err)

	t.Cleanup(func() {
		_, err = psqlUS.DeleteUser(uid)
		assert.Nil(t, err)
	})
}

func TestFindAllWorkoutsByUname(t *testing.T) {
	psqlUS, psqlWS, uid := setup(t)

	workout := entity.BaseWorkout{
		Username: "bobin",
		Title:    "Chest Day!",
	}

	count := 3
	for i := 0; i < count; i++ {
		newWorkout, err := psqlWS.Save(workout)
		assert.Nil(t, err)
		assert.NotEmpty(t, newWorkout)
		assert.Equal(t, "bobin", newWorkout.Username)
		assert.Equal(t, "Chest Day!", newWorkout.Title)

	}

	got, err := psqlWS.FindAllWorkoutsByUname("bobin")
	assert.Nil(t, err)
	assert.NotEmpty(t, got)
	assert.Equal(t, count, len(got))
	assert.Equal(t, "Chest Day!", got[0].BaseWorkout.Title)
	assert.Equal(t, "bobin", got[0].BaseWorkout.Username)

	t.Cleanup(func() {
		for i := 0; i < count; i++ {
			err := psqlWS.DeleteWorkout(got[i].ID)
			assert.Nil(t, err)
		}
		_, err = psqlUS.DeleteUser(uid)
		assert.Nil(t, err)
	})
}

// Util funcs.

func setup(t *testing.T) (us.UserStore, ws.WorkoutStore, int64) {
	psqlUS := psqlus.NewPsqlUserStore()
	user := entity.BaseUser{
		Username: "bobin",
		Email:    "bobin@gmail.com",
		Password: "bobinhood",
	}
	newUser, err := psqlUS.Save(user)

	assert.Nil(t, err)
	assert.NotEmpty(t, newUser)
	assert.Equal(t, "bobin", newUser.Username)

	psqlWS := psqlws.NewPsqlWorkoutStore()
	return psqlUS, psqlWS, newUser.ID
}
