package authtests

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/mhd53/quanta-fitness-server/internal/auth"
	us "github.com/mhd53/quanta-fitness-server/internal/datastore/userstore"
)

func TestValidateRegisterationMismatch(t *testing.T) {

	testValidator := auth.NewAuthValidator(nil)

	err := testValidator.ValidateRegisteration(MOCK_USERNAME, MOCK_EMAIL, MOCK_PWD, "hadio")

	assert.NotNil(t, err)
	assert.Equal(t, "Password must equal Confirm!", err.Error())
}

func TestValidateRegisterationUserExists(t *testing.T) {
	mockStore := us.NewMockUserStore()

	ucreated, _ := mockStore.Save(CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateRegisteration(MOCK_USERNAME, MOCK_EMAIL, MOCK_PWD, MOCK_PWD)

	assert.NotNil(t, err)
	assert.Equal(t, "Username already exists!", err.Error())
}

func TestValidateRegisterationEmailExists(t *testing.T) {
	mockStore := us.NewMockUserStore()

	ucreated, _ := mockStore.Save(CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateRegisteration("bobin", MOCK_EMAIL, MOCK_PWD, MOCK_PWD)

	assert.NotNil(t, err)
	assert.Equal(t, "Email already exists!", err.Error())
}

func TestValidateRegisterationWithInvalidEmail(t *testing.T) {
	mockStore := us.NewMockUserStore()
	notEmail := "notanemail.com"

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateRegisteration(MOCK_USERNAME, notEmail, MOCK_PWD, MOCK_PWD)

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid email address!", err.Error())
}

func TestValidateRegisterationSuccess(t *testing.T) {
	mockStore := us.NewMockUserStore()

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateRegisteration(MOCK_USERNAME, MOCK_EMAIL, MOCK_PWD, MOCK_PWD)

	assert.Nil(t, err)
}

func TestValidateLoginWithUnameWhenUserNotExist(t *testing.T) {
	mockStore := us.NewMockUserStore()

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateLoginWithUname(MOCK_USERNAME, MOCK_PWD)

	assert.NotNil(t, err)
	assert.Equal(t, "Username doesn't exist!", err.Error())
}

func TestValidateLoginWithUnameWhenIncorrectPwd(t *testing.T) {
	mockStore := us.NewMockUserStore()

	ucreated, _ := mockStore.Save(CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateLoginWithUname(MOCK_USERNAME, "bobin")

	assert.NotNil(t, err)
	assert.Equal(t, "Incorrect password!", err.Error())
}

func TestValidateLoginWithUnameSuccess(t *testing.T) {
	mockStore := us.NewMockUserStore()

	ucreated, _ := mockStore.Save(CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateLoginWithUname(MOCK_USERNAME, MOCK_PWD)

	assert.Nil(t, err)
}

func TestValidateLoginWithEmaileWhenUserNotExist(t *testing.T) {
	mockStore := us.NewMockUserStore()

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateLoginWithEmail(MOCK_EMAIL, MOCK_PWD)

	assert.NotNil(t, err)
	assert.Equal(t, "Email doesn't exist!", err.Error())
}

func TestValidateLoginWithEmailWhenIncorrectPwd(t *testing.T) {
	mockStore := us.NewMockUserStore()

	ucreated, _ := mockStore.Save(CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateLoginWithEmail(MOCK_EMAIL, "bobin")

	assert.NotNil(t, err)
	assert.Equal(t, "Incorrect password!", err.Error())
}

func TestValidateLoginWithEmailWhenInvalidEmail(t *testing.T) {
	mockStore := us.NewMockUserStore()

	ucreated, _ := mockStore.Save(CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateLoginWithEmail("notanemail.com", MOCK_PWD)

	assert.NotNil(t, err)
	assert.Equal(t, "Invalid email address!", err.Error())
}

func TestValidateLoginWithEmailSuccess(t *testing.T) {
	mockStore := us.NewMockUserStore()

	ucreated, _ := mockStore.Save(CreateValidAuthBaseUser())
	assert.NotEmpty(t, ucreated)

	testValidator := auth.NewAuthValidator(mockStore)

	err := testValidator.ValidateLoginWithEmail(MOCK_EMAIL, MOCK_PWD)

	assert.Nil(t, err)
}
