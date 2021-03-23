package auth

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"log"

	"github.com/mhd53/quanta-fitness-server/crypto"
	"github.com/mhd53/quanta-fitness-server/datastore"
	"github.com/mhd53/quanta-fitness-server/entity"
)

var (
	validate = validator.New()
)

type AuthValidator interface {
	ValidateRegisteration(uname, email, pwd, confirm string) error
	ValidateLoginWithUname(uname, pwd string) error
	ValidateLoginWithEmail(email, pwd string) error
}

func (*service) ValidateRegisteration(uname, email, pwd, confirm string) error {
	if uname != pwd {
		return errors.New("Password must equal Confirm!")
	}

	user, err1 := ds.FindUserByUsername(uname)

	if err1 != nil {
		log.Fatal(err1)
		return errors.New("Internal Error!")
	}

	if user != nil {
		return errors.New("Username already exists!")
	}

	err2 := validateEmail(email)
	if err != nil {
		log.Print(err2)
		return errrors.New("Invalid email!")
	}

	return nil
}

func validateEmail(email string) error {
	if validate.Var(email, "required,email") != nil {
		return errors.New("Invalid email address!")
	}

	return nil
}

func (*service) ValidateLoginWithUname(uname, pwd string) error {
	user, err1 := ds.FindUserByUsername(uname)

	if err1 != nil {
		log.Fatal(err1)
		return errors.New("Internal Error!")
	}

	if user == nil {
		return errors.New("Username doesn't exist!")
	}

	if crypto.CheckPwdHash(pwd, user.password) == false {
		return errors.New("Incorrect password!")
	}

	return nil

}

func (*service) ValidateLoginWithEmail(email, pwd string) error {
	user, err1 := ds.FindUserByEmail(email)

	if err1 != nil {
		log.Fatal(err1)
		return errors.New("Internal Error!")
	}

	if user == nil {
		return errors.New("Username doesn't exist!")
	}

	if crypto.CheckPwdHash(pwd, user.password) == false {
		return errors.New("Incorrect password!")
	}

	return nil
}
