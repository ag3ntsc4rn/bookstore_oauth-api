package accesstoken

import (
	"strings"

	"github.com/ag3ntsc4rn/bookstore_oauth-api/src/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(tokenID string) (*AccessToken, *errors.RestErr) {
	tokenID = strings.TrimSpace(tokenID)
	if len(tokenID) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.repository.GetById(tokenID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
