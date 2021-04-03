/*
This file contains implementation details for accessing a Postgres db for
workoutstore.

Created At: 2021/04/03
Author: mhd53
*/
package psqlws

import (
	"errors"
	"log"

	"github.com/mhd53/quanta-fitness-server/database/psql"
	ws "github.com/mhd53/quanta-fitness-server/internal/datastore/workoutstore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type store struct{}

func NewPsqlWorkoutStore() ws.WorkoutStore {
	return &store{}
}

func (*store) Save(workout entity.BaseWorkout) (entity.Workout, error) {
	db, err := psql.ConnectDB()
	if err != nil {
		log.Printf("Error %s: failed to connect to db!", err)
		return entity.Workout{}, err
	}
	defer db.Close()

	var uid int64
	var dbWorkout entity.Workout
	dbWorkout.BaseWorkout.Username = workout.Username

	query := `
	SELECT id FROM users
	WHERE username = $1
	`

	err = db.QueryRow(query, workout.Username).Scan(&uid)
	if err != nil {
		log.Printf("Error %s: failed get user_id from users table!", err)
		return entity.Workout{}, err
	}

	query = `
	INSERT INTO workouts(user_id, title)
	VALUES($1, $2)
	RETURNING id, title, created_at, updated_at
	`

	err = db.QueryRow(query, uid, workout.Title).Scan(&dbWorkout.ID, &dbWorkout.BaseWorkout.Title, &dbWorkout.CreatedAt, &dbWorkout.UpdatedAt)
	if err != nil {
		log.Printf("Error %s: couldn't insert new workout into db!", err)
		return entity.Workout{}, err
	}

	return dbWorkout, nil
}

func (*store) Update(wid int64, updates entity.BaseWorkout) error {
	return nil
}

func (*store) FindWorkoutById(wid int64) (entity.Workout, bool, error) {
	return entity.Workout{}, false, nil
}

func (*store) DeleteWorkout(wid int64) error {
	db, err := psql.ConnectDB()
	if err != nil {
		log.Printf("Error %s: failed to connect to db!", err)
		return err
	}
	defer db.Close()

	query := `
	DELETE FROM workouts
	WHERE id = $1;
	`

	res, err := db.Exec(query, wid)
	if err != nil {
		log.Printf("Error %s: couldn't delete workout from db!", err)
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s: couldn't get rows affected", err)
		return err
	}

	if count != 1 {
		return errors.New("The number of affected rows is not 1!")
	}

	return nil
}

func (*store) FindAllWorkoutsByUname(uname string) ([]entity.Workout, error) {
	var workouts []entity.Workout
	return workouts, nil
}
