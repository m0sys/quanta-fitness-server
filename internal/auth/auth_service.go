package auth

import (
	"errors"
	"log"

	"github.com/mhd53/quanta-fitness-server/internal/datastore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
	"github.com/mhd53/quanta-fitness-server/pkg/crypto"
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
		return "", err
	}

	token, err2 := crypto.GenerateToken(uname)

	if err2 != nil {
		log.Fatal(err2)
		return "", mainErr
	}

	user := entity.BaseUser{
		Username: uname,
		Email:    email,
		Password: pwd,
	}

	_, err3 := ds.Save(user)
	if err3 != nil {
		log.Fatal(err3)
		return "", mainErr
	}

	return token, nil
}

func (*service) LoginWithUname(uname, pwd string) (string, error) {
	mainErr := errors.New("Failed to log user in! Please try again later.")
	err := val.ValidateLoginWithUname(uname, pwd)

	if err != nil {
		return "", err
	}

	token, err2 := crypto.GenerateToken(uname)
	if err2 != nil {
		log.Fatal(err2)
		return "", mainErr
	}

	return token, nil
}

func (*service) LoginWithEmail(email, pwd string) (string, error) {
	mainErr := errors.New("Failed to log user in! Please try again later.")
	err := val.ValidateLoginWithEmail(email, pwd)

	if err != nil {
		return "", err
	}

	user, _, err2 := ds.FindUserByEmail(email)
	if err2 != nil {
		log.Fatal(err2)
		return "", mainErr
	}

	token, err3 := crypto.GenerateToken(user.Username)
	if err2 != nil {
		log.Fatal(err3)
		return "", mainErr
	}

	return token, nil
}

func (*service) Logout(string) error {
	return nil
}
