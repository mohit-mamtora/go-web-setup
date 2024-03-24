package repository

import (
	"github.com/jmoiron/sqlx"
	app "github.com/mohit-mamtora/go-web-setup/app"
)

type (
	Repository struct {
		user User
	}
)

func InitializeRepository(db *sqlx.DB, dh *app.DependencyHandler) *Repository {
	return &Repository{
		user: &UserRepository{
			DB:  db,
			log: dh.Logger,
		},
	}
}
