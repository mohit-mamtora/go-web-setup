package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/mohit-mamtora/go-web-setup/app/logger"
	usermodel "github.com/mohit-mamtora/go-web-setup/app/model"
	"github.com/mohit-mamtora/go-web-setup/app/model/dto"
)

type (
	User interface {
		GetById(ctx context.Context, id uuid.UUID) (*usermodel.User, error)
		Update(ctx context.Context, user *dto.UserProfileUpdate, id uuid.UUID) error
		Register(context.Context, *usermodel.User) error
		Login(ctx context.Context, username, password string) (*usermodel.User, error)
		DeleteToken(ctx context.Context, token *usermodel.Auth) error
		ValidateToken(ctx context.Context, token *usermodel.Auth) (bool, error)
		DeleteById(ctx context.Context, id uuid.UUID) error
		RegisterToken(ctx context.Context, auth *usermodel.Auth) error
	}

	UserRepository struct {
		DB  *sqlx.DB
		log logger.Log
	}
)

func (userRepository *UserRepository) DeleteById(ctx context.Context, id uuid.UUID) error {

	query, err := userRepository.DB.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	result, err := query.ExecContext(ctx, "delete from users_tokens where user_id=$1", id)

	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return errors.New("no users_tokens rows affected")
	}

	result, err = query.ExecContext(ctx, "delete from users where id=$1", id)

	if err != nil {
		query.Rollback()
		return err
	}

	rowAffected, err = result.RowsAffected()

	if err != nil {
		query.Rollback()
		return err
	}

	if rowAffected == 0 {
		query.Rollback()
		return errors.New("no users rows affected")
	}

	err = query.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (userRepository *UserRepository) DeleteToken(ctx context.Context, token *usermodel.Auth) error {
	result, err := userRepository.DB.ExecContext(ctx, "delete from users_tokens where id=$1 and user_id=$2", token.TokenId, token.UserId)

	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (userRepository *UserRepository) ValidateToken(ctx context.Context, token *usermodel.Auth) (bool, error) {
	fmt.Println(token)
	rows, err := userRepository.DB.QueryxContext(ctx, "select count(id) from users_tokens where id=$1 and user_id=$2", token.TokenId, token.UserId)

	if err != nil {
		return false, err
	}
	var count = new(int)
	rows.Next()

	err = rows.Scan(count)

	if err != nil {
		return false, err
	}

	fmt.Println(*count)

	return *count != 0, nil
}

func (userRepository *UserRepository) Update(ctx context.Context, user *dto.UserProfileUpdate, id uuid.UUID) error {

	_, err := userRepository.DB.ExecContext(ctx, "update users set email=$1, name=$2 where id = $3", user.Email, user.Name, id)
	if err != nil {
		return err
	}
	return nil
}

func (userRepository *UserRepository) GetById(ctx context.Context, id uuid.UUID) (*usermodel.User, error) {
	rows, err := userRepository.DB.QueryxContext(ctx, "select * from users where id = $1;", id)

	if err != nil {
		return nil, err
	}

	var user = usermodel.User{}
	rows.Next()
	err = rows.StructScan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepository *UserRepository) Register(ctx context.Context, user *usermodel.User) error {

	passwordHex, err := createHash(user.Password)

	if err != nil {
		return errors.New("error to create hash")
	}

	_, err = userRepository.DB.ExecContext(ctx, "INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4);", user.Id, user.Username, user.Email, passwordHex)

	if err != nil {
		pgError, ok := err.(*pq.Error)
		if !ok {
			return err
		}

		if pgError.Code == "23505" {
			return errors.New("duplicate data entry violation")
		}

		return err
	}

	return nil
}

func (userRepository *UserRepository) Login(ctx context.Context, username, password string) (*usermodel.User, error) {

	passwordHex, err := createHash(password)

	if err != nil {
		return nil, errors.New("error to create hash")
	}

	rows, err := userRepository.DB.QueryxContext(ctx, "select * from users where username = $1 and password = $2 limit 1;", username, passwordHex)

	if err != nil {
		return nil, err
	}

	var user = usermodel.User{}
	rows.Next()
	err = rows.StructScan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepository *UserRepository) RegisterToken(ctx context.Context, auth *usermodel.Auth) error {
	_, err := userRepository.DB.ExecContext(ctx, "INSERT INTO users_tokens (id, user_id) VALUES ($1, $2);", auth.TokenId, auth.UserId)

	if err != nil {
		pgError, ok := err.(*pq.Error)
		if !ok {
			return err
		}

		if pgError.Code == "23505" {
			return errors.New("duplicate data entry violation")
		}

		return err
	}
	return nil
}
