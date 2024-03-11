package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type (
	User struct {
		Id        uuid.UUID
		Name      string
		Username  string
		Password  string
		Balance   int64
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	Auth struct {
		Id       uuid.UUID
		Username string
		jwt.RegisteredClaims
	}
)
