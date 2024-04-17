package routes

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/mohit-mamtora/go-web-setup/app/model"
	"github.com/mohit-mamtora/go-web-setup/config"
)

func (route *Route) authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		err := jwtConfig(func(c echo.Context) error { return nil })(c)

		if err != nil {
			return echo.ErrUnauthorized
		}

		// verify token
		auth := getAuth(c)
		isValid, _ := route.service.UserService.ValidateToken(c.Request().Context(), auth)
		if !isValid {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

var jwtConfig = echojwt.WithConfig(echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(model.Auth)
	},
	SigningKey: []byte(config.AppKey),
})
