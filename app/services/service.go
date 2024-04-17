package services

import (
	"github.com/mohit-mamtora/go-web-setup/app"
	"github.com/mohit-mamtora/go-web-setup/app/repository"
)

type (
	Service struct {
		UserService UserService
	}
)

func InitializeService(repo *repository.Repository, dh *app.DependencyHandler) *Service {

	return &Service{
		UserService: &UserServiceImpl{
			Log:        dh.Logger,
			Repository: repo,
		},
	}
}
