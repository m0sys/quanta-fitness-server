package psqlustest

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mhd53/quanta-fitness-server/database/psql/psqlus"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

func TestSave(t *testing.T) {
	psqlDB := psqlus.NewPsqlUserStore()
	user := entity.BaseUser{
		Username: "robin",
		Email:    "robin@gmail.com",
		Password: "robinhood",
	}
	newUser, err := psqlDB.Save(user)

	assert.Nil(t, err)
	assert.NotEmpty(t, newUser)
	assert.Equal(t, "robin", newUser.Username)

	t.Cleanup(func() {
		_, err = psqlDB.DeleteUser(newUser.ID)
		assert.Nil(t, err)
	})
}

func TestFindUserByUsername(t *testing.T) {
	psqlDB := psqlus.NewPsqlUserStore()
	user := entity.BaseUser{
		Username: "robin",
		Email:    "robin@gmail.com",
		Password: "robinhood",
	}
	newUser, err := psqlDB.Save(user)

	assert.Nil(t, err)
	assert.NotEmpty(t, newUser)
	assert.Equal(t, "robin", newUser.Username)

	got, found, err := psqlDB.FindUserByUsername("robin")
	assert.Nil(t, err)
	assert.True(t, found)
	assert.Equal(t, "robin", got.Username)

	t.Cleanup(func() {
		_, err = psqlDB.DeleteUser(newUser.ID)
		assert.Nil(t, err)
	})
}

func TestFindUserByEmail(t *testing.T) {
	psqlDB := psqlus.NewPsqlUserStore()
	user := entity.BaseUser{
		Username: "robin",
		Email:    "robin@gmail.com",
		Password: "robinhood",
	}
	newUser, err := psqlDB.Save(user)

	assert.Nil(t, err)
	assert.NotEmpty(t, newUser)
	assert.Equal(t, "robin", newUser.Username)

	got, found, err := psqlDB.FindUserByEmail("robin@gmail.com")
	assert.Nil(t, err)
	assert.True(t, found)
	assert.Equal(t, "robin", got.Username)
	assert.Equal(t, "robin@gmail.com", got.Email)

	t.Cleanup(func() {
		_, err = psqlDB.DeleteUser(newUser.ID)
		assert.Nil(t, err)
	})
}

func TestDelete(t *testing.T) {
	psqlDB := psqlus.NewPsqlUserStore()
	user := entity.BaseUser{
		Username: "robin",
		Email:    "robin@gmail.com",
		Password: "robinhood",
	}
	newUser, err := psqlDB.Save(user)

	assert.Nil(t, err)
	assert.NotEmpty(t, newUser)
	assert.Equal(t, "robin", newUser.Username)

	success, err := psqlDB.DeleteUser(newUser.ID)
	assert.Nil(t, err)
	assert.True(t, success)
}
