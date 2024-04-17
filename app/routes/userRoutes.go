package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mohit-mamtora/go-web-setup/app/model/dto"
)

func (route *Route) UserGroupRoute() {
	userGroup := route.Echo.Group("/api/v1").Group("/user", route.authMiddleware)

	userGroup.GET("/profile", func(c echo.Context) error {
		auth := getAuth(c)
		response, err := route.service.UserService.Profile(c.Request().Context(), auth.UserId)
		if err != nil {
			return echo.ErrInternalServerError
		}
		return c.JSON(http.StatusOK, response)
	})

	userGroup.PUT("/update/profile", func(c echo.Context) error {

		request := new(dto.UserProfileUpdate)

		if err := c.Bind(request); err != nil {
			return echo.ErrInternalServerError
		}

		if err := c.Validate(request); err != nil {
			return err
		}

		auth := getAuth(c)
		response, err := route.service.UserService.Update(c.Request().Context(), request, auth.UserId)
		if err != nil {
			route.log.Error("%v", err)
			return echo.ErrInternalServerError
		}
		return c.JSON(http.StatusOK, response)
	})

	userGroup.DELETE("/de-register", func(c echo.Context) error {

		auth := getAuth(c)

		err := route.service.UserService.Delete(c.Request().Context(), auth.UserId)

		if err != nil {
			return echo.ErrInternalServerError
		}
		return c.JSON(http.StatusOK, map[string]string{
			"status": "successful",
		})
	})

	userGroup.POST("/logout", func(c echo.Context) error {

		auth := getAuth(c)

		err := route.service.UserService.Logout(c.Request().Context(), auth)

		if err != nil {
			return echo.ErrInternalServerError
		}
		return c.JSON(http.StatusOK, map[string]string{
			"status": "successful",
		})
	})
}
