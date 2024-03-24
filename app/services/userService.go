package services

import (
	"github.com/mohit-mamtora/go-web-setup/app/logger"
	"github.com/mohit-mamtora/go-web-setup/app/model/dto"
	"github.com/mohit-mamtora/go-web-setup/app/repository"
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
