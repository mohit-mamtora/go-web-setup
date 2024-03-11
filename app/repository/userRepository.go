package repository

import (
	"github.com/Mohit-Mamtora/gofinlop/app/logger"
	"github.com/jmoiron/sqlx"
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
