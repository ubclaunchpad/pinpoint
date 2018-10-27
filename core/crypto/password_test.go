package crypto

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	unhashed := "1234test"
	hashed, err := hashAndSalt(unhashed)
	assert.Nil(t, err)
	assert.NotEqual(t, unhashed, hashed)


}

func TestIsLegalPassword(t *testing.T) {
	password := "Hxt-f$T@4%7"
	assert.True(t,IsLegalPassword(password))

	password2 := "Hxt  f$T@4%7"
	assert.False(t,IsLegalPassword(password2))
}

func TestIsLegalUsername(t *testing.T) {
	username := "Robert"
	assert.True(t,IsLegalUserName(username))

	username2 := "R o b e r t"
	assert.False(t,IsLegalUserName(username2))

	username3 := "Robert "
	assert.False(t,IsLegalUserName(username3))

	username4 := "Robert@lp"
	assert.False(t,IsLegalUserName(username4))
}

func TestCorrectPassword(t *testing.T) {
	unhashed := "LaunchPad@pinpoint"
	hashed, err := hashAndSalt(unhashed)
	assert.Nil(t, err)
	assert.NotEqual(t, unhashed, hashed)

	correct := comparePasswords(hashed, unhashed)
	assert.True(t, correct)

	correct = comparePasswords(hashed, "LaunchPadpinpoint")
	assert.False(t, correct)
}

func TestValidateCredentialValues(t *testing.T) {
	err := ValidateCredentialValues("finasdfsdfe", "okaasdfasdy")
	assert.Nil(t, err)

	err = ValidateCredentialValues("0123456789a", "0123456789")
	assert.Nil(t, err)

	err = ValidateCredentialValues("mojave", "mojave")
	assert.Equal(t, errSameUsernamePassword, err)

	err = ValidateCredentialValues("wowwow", "oh")
	assert.Equal(t, errInvalidPassword, err)

	err = ValidateCredentialValues("um", "ohasdf")
	assert.Equal(t, errInvalidUsername, err)

	err = ValidateCredentialValues("SpiderMan!!!!!!", "oasdfasdfh")
	assert.Equal(t, errInvalidUsername, err)

	err = ValidateCredentialValues("wowwow", "oasdfasdfh!!!!")
	assert.Nil(t, err)
}