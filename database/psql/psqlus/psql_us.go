/*
This file contains implementation details for accessing a Postgres db for
userstore.

Created At: 2021/04/03
Author: mhd53
*/
package psqlus

import (
	"log"

	"github.com/mhd53/quanta-fitness-server/database/psql"
	us "github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
)

type store struct{}

func NewPsqlUserStore() us.UserStore {
	return &store{}
}

func (*store) Save(user entity.BaseUser) (entity.User, error) {
	// Connect to db.
	db, err := psql.ConnectDB()
	if err != nil {
		log.Printf("Error %s: failed to connect to db!", err)
		return entity.User{}, err
	}
	defer db.Close()

	query := `
	INSERT INTO users(username, email, hashed_password)
	VALUES($1, $2, $3)
	RETURNING *
	`
	var dbUser entity.User

	err = db.QueryRow(query, user.Username, user.Email, user.Password).Scan(&dbUser.ID, &dbUser.BaseUser.Username, &dbUser.BaseUser.Email, &dbUser.BaseUser.Password, &dbUser.Joined, &dbUser.UpdatedAt, &dbUser.Weight, &dbUser.Height, &dbUser.Gender)
	if err != nil {
		log.Printf("Error %s: couldn't insert new user into db!", err)
		return entity.User{}, err
	}

	return dbUser, nil
}

func (*store) FindUserByUsername(username string) (entity.User, bool, error) {
	db, err := psql.ConnectDB()
	if err != nil {
		log.Printf("Error %s: failed to connect to db!", err)
		return entity.User{}, false, err
	}
	defer db.Close()

	query := `
	SELECT * FROM users
	WHERE username = $1
	`
	var dbUser entity.User

	err = db.QueryRow(query, username).Scan(&dbUser.ID, &dbUser.BaseUser.Username, &dbUser.BaseUser.Email, &dbUser.BaseUser.Password, &dbUser.Joined, &dbUser.UpdatedAt, &dbUser.Weight, &dbUser.Height, &dbUser.Gender)
	if err != nil {
		log.Printf("Error %s: couldn't query users table via uname!", err)
		return entity.User{}, false, err
	}

	return dbUser, true, nil

}

func (*store) FindUserByEmail(email string) (entity.User, bool, error) {
	db, err := psql.ConnectDB()
	if err != nil {
		log.Printf("Error %s: failed to connect to db!", err)
		return entity.User{}, false, err
	}
	defer db.Close()

	query := `
	SELECT * FROM users
	WHERE email = $1
	`
	var dbUser entity.User

	err = db.QueryRow(query, email).Scan(&dbUser.ID, &dbUser.BaseUser.Username, &dbUser.BaseUser.Email, &dbUser.BaseUser.Password, &dbUser.Joined, &dbUser.UpdatedAt, &dbUser.Weight, &dbUser.Height, &dbUser.Gender)
	if err != nil {
		log.Printf("Error %s: couldn't query users table via email!", err)
		return entity.User{}, false, err
	}

	return dbUser, true, nil

}

func (*store) DeleteUser(id int64) (bool, error) {
	// Connect to db.
	db, err := psql.ConnectDB()
	if err != nil {
		log.Printf("Error %s: failed to connect to db!", err)
		return false, err
	}
	defer db.Close()

	query := `
	DELETE FROM users
	WHERE id = $1;
	`

	res, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error %s: couldn't delete user from db!", err)
		return false, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s: couldn't get rows affected", err)
		return false, err
	}

	return count == 1, nil
}
