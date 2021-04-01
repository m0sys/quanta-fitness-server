package auth

import (
	"fmt"

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
		return entity.User{}, err
	}

	hashedPwd, err3 := crypto.HashPwd(pwd)
	if err3 != nil {
		return entity.User{}, formatErr(err3)
	}

	user := entity.BaseUser{
		Username: uname,
		Email:    email,
		Password: hashedPwd,
	}

	userDS, err3 := ds.Save(user)
	if err3 != nil {
		return entity.User{}, formatErr(err3)
	}

	return userDS, nil
}

func (*service) LoginWithUname(uname, pwd string) (entity.User, error) {
	return login(uname, pwd, val.ValidateLoginWithUname, ds.FindUserByUsername)
}

func (*service) LoginWithEmail(email, pwd string) (entity.User, error) {
	return login(email, pwd, val.ValidateLoginWithEmail, ds.FindUserByEmail)
}

func login(cred, pwd string, validator func(cred, pwd string) error, fetcher func(cred string) (entity.User, bool, error)) (entity.User, error) {
	err := validator(cred, pwd)

	if err != nil {
		return entity.User{}, err
	}

	userDS, _, err2 := fetcher(cred)
	if err2 != nil {
		return entity.User{}, formatErr(err2)
	}

	return userDS, nil
}

func (*service) Logout(string) error {
	return nil
}

func formatErr(err error) error {
	return fmt.Errorf("%s: couldn't access db: %w", "auth_service", err)
}
