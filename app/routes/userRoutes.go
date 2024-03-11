package routes

import "github.com/labstack/echo/v4"

func (route *Route) UserGroupRoute() {
	userGroup := route.Echo.Group("api/v1").Group("/user")

	userGroup.GET("/profile", func(c echo.Context) error {
		return nil
	})
}
