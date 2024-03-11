package services

import (
	"github.com/Mohit-Mamtora/gofinlop/app"
	"github.com/Mohit-Mamtora/gofinlop/app/repository"
)

type (
	Service struct {
		UserService UserService
	}
)

func InitilizeService(repo *repository.Repository, dh *app.DependencyHandler) *Service {
	return &Service{
		UserService: &UserServiceImpl{
			Log:        dh.Logger,
			Repository: repo,
		},
	}
}
