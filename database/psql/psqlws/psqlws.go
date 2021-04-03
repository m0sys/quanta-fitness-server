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
	db, err := psql.ConnectDB()
	if err != nil {
		log.Printf("Error %s: failed to connect to db!", err)
		return err
	}
	defer db.Close()

	query := `
	UPDATE workouts
	SET title = $2
	WHERE id = $1;
	`

	res, err := db.Exec(query, wid, updates.Title)
	if err != nil {
		log.Printf("Error %s: couldn't update workout from db!", err)
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

func (*store) FindWorkoutById(wid int64) (entity.Workout, bool, error) {
	db, err := psql.ConnectDB()
	if err != nil {
		log.Printf("Error %s: failed to connect to db!", err)
		return entity.Workout{}, false, err
	}
	defer db.Close()

	var dbWorkout entity.Workout

	query := `
	SELECT w.id, w.title, w.created_at, w.updated_at, u.username 
	FROM workouts AS w
 	JOIN users AS u ON w.user_id = u.id
	WHERE w.id = $1;
	`

	err = db.QueryRow(query, wid).Scan(&dbWorkout.ID, &dbWorkout.BaseWorkout.Title, &dbWorkout.CreatedAt, &dbWorkout.UpdatedAt, &dbWorkout.BaseWorkout.Username)
	if err != nil {
		log.Printf("Error %s: couldn't find workout db!", err)
		return entity.Workout{}, false, err
	}

	return dbWorkout, true, nil
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

	db, err := psql.ConnectDB()
	if err != nil {
		log.Printf("Error %s: failed to connect to db!", err)
		return workouts, err
	}
	defer db.Close()

	var uid int64

	query := `
	SELECT id FROM users
	WHERE username = $1
	`

	err = db.QueryRow(query, uname).Scan(&uid)
	if err != nil {
		log.Printf("Error %s: failed get user_id from users table!", err)
		return workouts, err
	}

	query = `SELECT id, title, created_at, updated_at FROM workouts WHERE user_id = $1;`

	rows, err := db.Query(query, uid)
	if err != nil {
		log.Printf("Error %s: couldn't get rows from workout table!", err)
		return workouts, err

	}
	defer rows.Close()

	for rows.Next() {
		var workout entity.Workout
		workout.BaseWorkout.Username = uname

		err = rows.Scan(&workout.ID, &workout.BaseWorkout.Title, &workout.CreatedAt, &workout.UpdatedAt)
		if err != nil {
			log.Printf("Error %s: couldn't scan row from workout table!", err)
			return workouts, err

		}

		workouts = append(workouts, workout)
	}

	err = rows.Err()
	if err != nil {
		log.Printf("Error %s: while looping through workouts table!", err)
		return workouts, err
	}

	return workouts, nil
}
