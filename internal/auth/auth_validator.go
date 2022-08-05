package auth

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"

	"github.com/m0sys/quanta-fitness-server/internal/datastore/userstore"
	"github.com/m0sys/quanta-fitness-server/pkg/crypto"
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

	_, found, err := valStore.FindUserByUsername(uname)
	if err != nil {
		return fmt.Errorf("couldn't access db: %w", err)
	}

	if found {
		return errors.New("Username already exists!")
	}

	err2 := validateEmail(email)
	if err2 != nil {
		return err2
	}

	_, found2, err3 := valStore.FindUserByEmail(email)
	if err3 != nil {
		return fmt.Errorf("couldn't access db: %w", err3)
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
	user, found, err := valStore.FindUserByUsername(uname)
	if err != nil {
		return fmt.Errorf("couldn't access db: %w", err)
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
		return err
	}

	user, found, err := valStore.FindUserByEmail(email)
	if err != nil {
		return fmt.Errorf("couldn't access db: %w", err)
	}

	if !found {
		return errors.New("Email doesn't exist!")
	}

	if !crypto.CheckPwdHash(pwd, user.Password) {
		return errors.New("Incorrect password!")
	}

	return nil
}
