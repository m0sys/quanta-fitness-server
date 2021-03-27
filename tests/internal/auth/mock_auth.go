package authtests

import (
	"github.com/mhd53/quanta-fitness-server/internal/entity"
	"github.com/mhd53/quanta-fitness-server/pkg/crypto"
)

const (
	MOCK_ID       int64   = 1
	MOCK_USERNAME         = "robin"
	MOCK_EMAIL            = "robin@gmail.com"
	MOCK_PWD              = "robinhood"
	MOCK_WEIGHT   float32 = 75.0
	MOCK_HEIGHT   float32 = 158.0
	MOCK_GENDER           = "male"
)

func CreateValidMockUser(id int64) entity.User {
	hashedPwd, _ := crypto.HashPwd(MOCK_PWD)

	user := entity.User{
		ID: id,
		BaseUser: entity.BaseUser{
			Username: MOCK_USERNAME,
			Email:    MOCK_EMAIL,
			Password: hashedPwd,
		},
		Weight: MOCK_WEIGHT,
		Height: MOCK_HEIGHT,
		Gender: MOCK_GENDER,
	}

	return user
}

func CreateValidAuthBaseUser() entity.BaseUser {
	hashedPwd, _ := crypto.HashPwd(MOCK_PWD)

	user := entity.BaseUser{
		Username: MOCK_USERNAME,
		Email:    MOCK_EMAIL,
		Password: hashedPwd,
	}

	return user
}
