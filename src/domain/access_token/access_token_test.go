package accesstoken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}
func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.NotNil(t, at, "brand new access token should not be nil")
	assert.EqualValues(t, false, at.IsExpired(), "brand new access token should not be expired")
	assert.EqualValues(t, "", at.AccessToken, "brand new access token should not have a token id")
	assert.EqualValues(t, 0, at.UserID, "brand new access token should not have associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.EqualValues(t, true, at.IsExpired(), "brand new access token should be expired by default.")
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.EqualValues(t, false, at.IsExpired(), "brand new access token with expiration 3 hours from now should not be expired.")
}
