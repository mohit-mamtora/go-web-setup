package services

import (
	"github.com/Mohit-Mamtora/gofinlop/app/logger"
	"github.com/Mohit-Mamtora/gofinlop/app/model/dto"
	"github.com/Mohit-Mamtora/gofinlop/app/repository"
)

type (
	UserService interface {
		Login(*dto.Request) (*dto.Response, error)
	}

	UserServiceImpl struct {
		Log        logger.Log
		Repository *repository.Repository
	}
)

func (userService *UserServiceImpl) Login(req *dto.Request) (*dto.Response, error) {
	return &dto.Response{
		"status": "successful",
	}, nil
}
