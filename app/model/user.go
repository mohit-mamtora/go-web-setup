package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type (
	User struct {
		Id        uuid.UUID `json:"id"`
		Name      string    `json:"name" validate:"required"`
		Email     string    `json:"email" validate:"required"`
		Username  string    `json:"username" validate:"required"`
		Password  string    `json:"-"`
		CreatedAt time.Time `json:"created_at" db:"created_at"`
		UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	}

	Auth struct {
		TokenId   uuid.UUID `json:"token_id"  db:"id"`
		UserId    uuid.UUID `json:"user_id"`
		Username  string    `json:"username"`
		CreatedAt time.Time `json:"created_at" db:"created_at"`
		UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
		jwt.RegisteredClaims
	}
)
