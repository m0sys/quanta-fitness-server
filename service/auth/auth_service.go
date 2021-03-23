package auth

import (
	"errors"
	"log"

	"github.com/mhd53/quanta-fitness-server/crypto"
	"github.com/mhd53/quanta-fitness-server/datastore"
	"github.com/mhd53/quanta-fitness-server/entity"
)

type AuthService interface {
	Register(uname, email, pwd, confirm string) (string, error)
	LoginWithUname(uname, pwd string) (string, error)
	LoginWithEmail(email, pwd string) (string, error)
	Logout(token string) error
}

type service struct{}

var (
	ds  datastore.UserStore
	val AuthValidator
)

func NewAuthService(datastore datastore.UserStore, validator AuthValidator) AuthService {
	ds = datastore
	val = validator
	return &service{}
}

func (*service) Register(uname, email, pwd, confirm string) (string, error) {
	mainErr := errors.New("Failed to register user! Please try again later.")

	err := val.ValidateRegisteration(uname, email, pwd, confirm)

	if err != nil {
		log.Print(err)
		return nil,  err
	}

	token, err2 := crypto.GenerateToken(user.Username)

	if err2 != nil {
		log.Fatal(err2)
		return nil, mainErr
	}

	_, err3 := ds.Save(user)
	if err3 != nil {
		log.Fatal(err3)
		return nil, mainErr
	}

	return token, nil
}

func (*service) LoginWithUname(uname, pwd string) (string, error) {
	mainErr := errors.New("Failed to log user in! Please try again later.")
	err := val.ValidateLoginWithUname(uname, pwd)

	if err != nil {
		return nil, err
	}

	token. err2 := crypto.GenerateToken(uname)
	if err2 != nil {
		log.Fatal(err2)
		return nil, mainErr
	}

	return token, nil
}

func (*service) LoginWithEmail(email, pwd string) (string, error) {
	mainErr := errors.New("Failed to log user in! Please try again later.")
	err := val.ValidateLoginWithEmail(uname, pwd)

	if err != nil {
		return nil, err
	}

	user, err2 := ds.FindUserByEmail(email)
	if err2 != nil {
		log.Fatal(err2)
		return nil, mainErr
	}

	token. err3 := crypto.GenerateToken(user.username)
	if err2 != nil {
		log.Fatal(err3)
		return nil, mainErr
	}

	return token, nil
}

func (*service) Logout(string) error {
	return nil
}
