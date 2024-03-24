package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mohit-mamtora/go-web-setup/app/model/dto"
)

func (route *Route) AuthRoutes() {

	/* LOGIN */
	route.Echo.POST("/login", func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")
		route.log.Info("login: %s=%s", username, password)
		response, err := route.service.UserService.Login(&dto.Request{
			"username": username,
			"password": password,
		})
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, response)
	})
}
