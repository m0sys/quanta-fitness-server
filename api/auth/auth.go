package auth

import (
	"errors"
	"log"

	serv "github.com/mhd53/quanta-fitness-server/internal/auth"
	ustore "github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
	"github.com/mhd53/quanta-fitness-server/pkg/crypto"
)

var (
	us        ustore.UserStore
	validator serv.AuthValidator
	service   serv.AuthService
	// internalErr = errors.New("Internal Error!")
	notImplErr = errors.New("Not Implemented!")
)

type ServerAuth interface {
	RegisterNewUser(uname, email, pwd, confirm string) (string, error)
	LoginWithUname(uname, pwd string) (string, error)
	LoginWithEmail(email, pwd string) (string, error)
}

type server struct{}

func NewServerAuth() ServerAuth {
	us = ustore.NewMockUserStore()
	validator = serv.NewAuthValidator(us)
	service = serv.NewAuthService(us, validator)
	return &server{}
}

func (*server) RegisterNewUser(uname, email, pwd, confirm string) (string, error) {
	user, err := service.Register(uname, email, pwd, confirm)

	if err != nil {
		log.Printf("Server Error: %s", err.Error())
		return "", err
	}

	token, err2 := crypto.GenerateToken(user.Username)

	if err2 != nil {
		log.Printf("Server Error: %s", err2.Error())
		return "", err2
	}

	return token, nil
}

func (*server) LoginWithUname(uname, pwd string) (string, error) {
	user, err := service.LoginWithUname(uname, pwd)

	if err != nil {
		log.Printf("Server Error: %s", err.Error())
		return "", err
	}

	token, err2 := crypto.GenerateToken(user.Username)

	if err2 != nil {
		log.Printf("Server Error: %s", err2.Error())
		return "", err2
	}

	return token, nil
}

func (*server) LoginWithEmail(email, pwd string) (string, error) {
	user, err := service.LoginWithEmail(email, pwd)

	if err != nil {
		log.Printf("Server Error: %s", err.Error())
		return "", err
	}

	token, err2 := crypto.GenerateToken(user.Username)

	if err2 != nil {
		log.Printf("Server Error: %s", err2.Error())
		return "", err2
	}

	return token, nil
}
