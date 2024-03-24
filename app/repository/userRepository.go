package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mohit-mamtora/go-web-setup/app/logger"
)

type (
	User interface {
		All()
	}

	UserRepository struct {
		DB  *sqlx.DB
		log logger.Log
	}
)

func (userRepository *UserRepository) All() {

}
