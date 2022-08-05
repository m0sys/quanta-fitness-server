package accounting

import (
	"log"

	"github.com/m0sys/quanta-fitness-server/account/user"
)

type AccountService struct {
	repo Repository
}

func NewAccountService(repository Repository) AccountService {
	return AccountService{repo: repository}
}

func (a AccountService) SignUp(uname, email, pwd string) (user.User, error) {
	if err := a.validateUname(uname); err != nil {
		return user.User{}, err
	}

	if err := a.validateEmail(email); err != nil {
		return user.User{}, err
	}

	newUser, err := user.NewUser(uname, email, pwd)
	if err != nil {
		return user.User{}, err
	}

	err = a.repo.StoreUser(newUser)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return user.User{}, errInternal
	}

	return newUser, nil
}

func (a AccountService) validateUname(uname string) error {
	_, found, err := a.repo.FindUserByUname(uname)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return errInternal
	}

	if found {
		return ErrUnameAlreadyExists
	}

	return nil
}

func (a AccountService) validateEmail(email string) error {
	_, found, err := a.repo.FindUserByEmail(email)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return errInternal
	}

	if found {
		return ErrEmailAlreadyExists
	}

	return nil
}

func (a AccountService) Login(uname, pwd string) (user.User, error) {
	userFound, found, err := a.repo.FindUserByUname(uname)
	if err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return user.User{}, errInternal
	}

	if !found {
		return user.User{}, ErrUnameNotFound
	}

	if err := userFound.CheckPassword(pwd); err != nil {
		log.Printf("%s: %s", errSlug, err.Error())
		return user.User{}, ErrIncorrectPassword
	}

	return userFound, nil
}
