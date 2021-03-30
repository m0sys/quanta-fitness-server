package auth

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"log"

	"github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
	"github.com/mhd53/quanta-fitness-server/pkg/crypto"
)

var (
	validate = validator.New()
	valStore userstore.UserStore
)

type authValidator struct{}

type AuthValidator interface {
	ValidateRegisteration(uname, email, pwd, confirm string) error
	ValidateLoginWithUname(uname, pwd string) error
	ValidateLoginWithEmail(email, pwd string) error
}

func NewAuthValidator(userStore userstore.UserStore) AuthValidator {
	valStore = userStore
	return &authValidator{}
}

func (*authValidator) ValidateRegisteration(uname, email, pwd, confirm string) error {
	if pwd != confirm {
		return errors.New("Password must equal Confirm!")
	}

	_, found, err1 := valStore.FindUserByUsername(uname)

	if err1 != nil {
		log.Fatal(err1)
		return errors.New("Internal Error!")
	}

	if found {
		return errors.New("Username already exists!")
	}

	err2 := validateEmail(email)
	if err2 != nil {
		log.Print(err2)
		return errors.New("Invalid email!")
	}

	_, found2, err3 := valStore.FindUserByEmail(email)
	if err3 != nil {
		log.Fatal(err3)
		return errors.New("Internal Error!")
	}

	if found2 {
		return errors.New("Email already exists!")
	}

	return nil
}

func validateEmail(email string) error {
	if validate.Var(email, "required,email") != nil {
		return errors.New("Invalid email address!")
	}

	return nil
}

func (*authValidator) ValidateLoginWithUname(uname, pwd string) error {
	user, found, err1 := valStore.FindUserByUsername(uname)

	if err1 != nil {
		log.Fatal(err1)
		return errors.New("Internal Error!")
	}

	if !found {
		return errors.New("Username doesn't exist!")
	}

	if !crypto.CheckPwdHash(pwd, user.Password) {
		return errors.New("Incorrect password!")
	}

	return nil

}

func (*authValidator) ValidateLoginWithEmail(email, pwd string) error {
	err := validateEmail(email)
	if err != nil {
		log.Print(err)
		return errors.New("Invalid email!")
	}

	user, found, err1 := valStore.FindUserByEmail(email)

	if err1 != nil {
		log.Fatal(err1)
		return errors.New("Internal Error!")
	}

	if !found {
		return errors.New("Email doesn't exist!")
	}

	if !crypto.CheckPwdHash(pwd, user.Password) {
		return errors.New("Incorrect password!")
	}

	return nil
}
