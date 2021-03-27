package util

import (
	"errors"
	"log"

	us "github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
)

func CheckUserExists(aus us.UserStore, uname string) (bool, error) {
	_, found, err := aus.FindUserByUsername(uname)
	if err != nil {
		log.Fatal(err)
		return false, errors.New("Internal error! Please try again later.")
	}

	if !found {
		return false, nil
	}

	return true, nil

}
