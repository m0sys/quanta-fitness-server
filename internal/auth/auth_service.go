package auth

import (
	"errors"
	"log"

	"github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
	"github.com/mhd53/quanta-fitness-server/internal/entity"
	"github.com/mhd53/quanta-fitness-server/pkg/crypto"
)

type AuthService interface {
	Register(uname, email, pwd, confirm string) (entity.User, error)
	LoginWithUname(uname, pwd string) (entity.User, error)
	LoginWithEmail(email, pwd string) (entity.User, error)
	Logout(token string) error
}

type service struct{}

var (
	ds  userstore.UserStore
	val AuthValidator
)

func NewAuthService(userstore userstore.UserStore, validator AuthValidator) AuthService {
	ds = userstore
	val = validator
	return &service{}
}

func (*service) Register(uname, email, pwd, confirm string) (entity.User, error) {

	err := val.ValidateRegisteration(uname, email, pwd, confirm)

	if err != nil {
		log.Print(err)
		return entity.User{}, err
	}

	hashedPwd, err3 := crypto.HashPwd(pwd)

	if err3 != nil {
		log.Fatal(err3)
		return entity.User{}, errors.New("Failed to register user! Please try again later.")
	}

	user := entity.BaseUser{
		Username: uname,
		Email:    email,
		Password: hashedPwd,
	}

	userDS, err4 := ds.Save(user)
	if err4 != nil {
		log.Fatal(err4)
		return entity.User{}, errors.New("Failed to register user! Please try again later.")
	}

	return userDS, nil
}

func (*service) LoginWithUname(uname, pwd string) (entity.User, error) {
	err := val.ValidateLoginWithUname(uname, pwd)

	if err != nil {
		return entity.User{}, err
	}

	userDS, _, err2 := ds.FindUserByUsername(uname)
	if err2 != nil {
		log.Fatal(err2)
		return entity.User{}, errors.New("Failed to login! Please try again later.")
	}

	return userDS, nil
}

func (*service) LoginWithEmail(email, pwd string) (entity.User, error) {
	err := val.ValidateLoginWithEmail(email, pwd)

	if err != nil {
		return entity.User{}, err
	}

	userDS, _, err2 := ds.FindUserByEmail(email)
	if err2 != nil {
		log.Fatal(err2)
		return entity.User{}, errors.New("Failed to register user! Please try again later.")
	}

	return userDS, nil
}

func (*service) Logout(string) error {
	return nil
}
