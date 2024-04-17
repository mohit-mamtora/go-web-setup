package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mohit-mamtora/go-web-setup/app/model/dto"
)

func (route *Route) AuthRoutes() {

	/* LOGIN */
	route.Echo.POST("/login", func(c echo.Context) error {

		request := new(dto.LoginRequest)

		if err := c.Bind(request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(request); err != nil {
			return err
		}

		response, err := route.service.UserService.Login(c.Request().Context(), request)

		if err != nil {
			route.log.Error("login Error: %v", err)
			return echo.ErrUnauthorized
		}

		return c.JSON(http.StatusOK, response)
	})

	/* REGISTER */
	route.Echo.POST("/register", func(c echo.Context) error {

		request := new(dto.RegisterRequest)

		if err := c.Bind(request); err != nil {
			route.log.Error("register Error: %v", err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(request); err != nil {
			route.log.Error("register Error: %v", err)
			return err
		}

		response, err := route.service.UserService.Register(c.Request().Context(), request)

		if err != nil {
			route.log.Error("register Error: %v", err)
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, response)
	})
}
