package db

import (
	accesstoken "github.com/ag3ntsc4rn/bookstore_oauth-api/src/domain/access_token"
	"github.com/ag3ntsc4rn/bookstore_oauth-api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*accesstoken.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (db *dbRepository) GetById(string) (*accesstoken.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection not implemented")
}

func NewRepository() DbRepository {
	return &dbRepository{}
}
