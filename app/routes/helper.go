package routes

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/mohit-mamtora/go-web-setup/app/model"
)

func getAuth(c echo.Context) *model.Auth {
	user := c.Get("user").(*jwt.Token)
	return user.Claims.(*model.Auth)
}
