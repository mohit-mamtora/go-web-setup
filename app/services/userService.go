package services

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mohit-mamtora/go-web-setup/app/logger"
	"github.com/mohit-mamtora/go-web-setup/app/model"
	"github.com/mohit-mamtora/go-web-setup/app/model/dto"
	"github.com/mohit-mamtora/go-web-setup/app/repository"
	"github.com/mohit-mamtora/go-web-setup/config"
)

type (
	UserService interface {
		Profile(context.Context, uuid.UUID) (*dto.Response, error)
		Update(context.Context, *dto.UserProfileUpdate, uuid.UUID) (*dto.Response, error)
		Login(context.Context, *dto.LoginRequest) (*dto.Response, error)
		Register(context.Context, *dto.RegisterRequest) (*dto.Response, error)
		ValidateToken(context.Context, *model.Auth) (bool, error)
		Logout(context.Context, *model.Auth) error
		Delete(context.Context, uuid.UUID) error
	}

	UserServiceImpl struct {
		Log        logger.Log
		Repository *repository.Repository
	}
)

func (userService *UserServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	err := userService.Repository.User.DeleteById(ctx, id)
	if err != nil {
		userService.Log.Info(err.Error())
		return err
	}
	return nil
}

func (userService *UserServiceImpl) Logout(ctx context.Context, token *model.Auth) error {
	err := userService.Repository.User.DeleteToken(ctx, token)
	if err != nil {
		userService.Log.Info(err.Error())
		return err
	}
	return nil
}

func (userService *UserServiceImpl) ValidateToken(ctx context.Context, token *model.Auth) (bool, error) {
	isValid, err := userService.Repository.User.ValidateToken(ctx, token)
	if err != nil {
		userService.Log.Info(err.Error())
		return false, err
	}
	return isValid, nil
}

func (userService *UserServiceImpl) Profile(ctx context.Context, id uuid.UUID) (*dto.Response, error) {
	user, err := userService.Repository.User.GetById(ctx, id)
	if err != nil {
		userService.Log.Info(err.Error())
		return nil, err
	}
	return &dto.Response{
		"data": user,
	}, nil
}

func (userService *UserServiceImpl) Update(ctx context.Context, request *dto.UserProfileUpdate, userId uuid.UUID) (*dto.Response, error) {

	err := userService.Repository.User.Update(ctx, request, userId)
	if err != nil {
		return nil, err
	}

	updatedUser, err := userService.Repository.User.GetById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &dto.Response{
		"user": updatedUser,
	}, nil
}

func (userService *UserServiceImpl) Login(ctx context.Context, req *dto.LoginRequest) (*dto.Response, error) {

	user, err := userService.Repository.User.Login(ctx, req.Username, req.Password)
	if err != nil {
		userService.Log.Info(err.Error())
		return nil, err
	}
	auth := model.Auth{
		TokenId:  uuid.New(),
		UserId:   user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, auth)
	token, err := jwtToken.SignedString([]byte(config.AppKey))
	if err != nil {
		return nil, err
	}

	err = userService.Repository.User.RegisterToken(ctx, &auth)
	if err != nil {
		userService.Log.Info(err.Error())
		return nil, err
	}

	return &dto.Response{
		"token": token,
	}, nil
}

func (userService *UserServiceImpl) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.Response, error) {

	err := userService.Repository.User.Register(ctx, &model.User{
		Id:       uuid.New(),
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})

	if err != nil {
		return nil, err
	}

	return &dto.Response{
		"status": "successful",
	}, nil
}
