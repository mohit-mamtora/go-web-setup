package repository

import (
	app "github.com/Mohit-Mamtora/gofinlop/app"
	"github.com/jmoiron/sqlx"
)

type (
	Repository struct {
		user User
	}
)

func InitilizeRepository(db *sqlx.DB, dh *app.DependencyHandler) *Repository {
	return &Repository{
		user: &UserRepository{
			DB:  db,
			log: dh.Logger,
		},
	}
}
