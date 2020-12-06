package accesstoken

import (
	"time"

	_"github.com/ag3ntsc4rn/bookstore_oauth-api/src/utils/errors"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (accessToken *AccessToken) IsExpired() bool {
	return time.Unix(accessToken.Expires, 0).Before(time.Now().UTC())
}
