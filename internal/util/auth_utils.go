package util

import (
	"fmt"
	us "github.com/m0sys/quanta-fitness-server/internal/datastore/userstore"
)

func CheckUserExists(aus us.UserStore, uname string) (bool, error) {
	_, found, err := aus.FindUserByUsername(uname)
	if err != nil {
		return false, fmt.Errorf("couldn't access db: %w", err)
	}

	if !found {
		return false, nil
	}

	return true, nil
}
