// User contains the User entity.
package user

import (
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var (
	validate = validator.New()
)

type User struct {
	username  string
	hashedPwd string
	email     string
	joined    time.Time
}

func NewUser(username, email, pwd string) (User, error) {
	if err := validateFields(username, email, pwd); err != nil {
		return User{}, err
	}

	hashedPwd, err := hashPassword(pwd)
	if err != nil {
		return User{}, err
	}

	return User{
		username:  username,
		hashedPwd: hashedPwd,
		email:     email,
		joined:    time.Now(),
	}, nil
}

// This should only be used to restore User with valid fields from persistence layer .
func RestoreUser(username, email, hashedPwd string, joined time.Time) (User, error) {
	if err := validateFields(username, email, hashedPwd); err != nil {
		return User{}, err
	}

	return User{
		username:  username,
		hashedPwd: hashedPwd,
		email:     email,
		joined:    joined,
	}, nil
}

func (u *User) Username() string {
	return u.username
}

func (u *User) HashedPwd() string {
	return u.hashedPwd
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Joined() time.Time {
	return u.joined
}

// Field validation.

func validateFields(uname, email, pwd string) error {
	if err := validateUname(uname); err != nil {
		return err
	}

	if err := validateEmail(email); err != nil {
		return err
	}

	if err := validatePassword(pwd); err != nil {
		return err
	}

	return nil
}

func validatePassword(pwd string) error {
	length := len(pwd)
	if length < 12 || length > 128 {
		return errInvalidPassword
	}
	return nil
}

func validateUname(uname string) error {
	if len(uname) < 3 {
		return errInvalidUname
	}
	return nil
}

func validateEmail(email string) error {
	if validate.Var(email, "required,email") != nil {
		return errInvalidEmail
	}

	return nil
}

// Encryption functions.

func hashPassword(pwd string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return "", errHashFailed
	}

	return string(hashedPwd), nil
}

func checkPassword(pwd string, hashedPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
}
