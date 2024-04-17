package routes

import (
	"context"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mohit-mamtora/go-web-setup/app"
	"github.com/mohit-mamtora/go-web-setup/app/logger"
	"github.com/mohit-mamtora/go-web-setup/app/services"
	"github.com/mohit-mamtora/go-web-setup/config"
)

type (
	Route struct {
		*echo.Echo
		log     logger.Log
		service *services.Service
	}
)

func InitializeRoute(service *services.Service, dh *app.DependencyHandler) *Route {
	e := echo.New()
	e.HideBanner = true
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Recover())

	// TODO
	if config.IsDebugMode {
		e.Use(middleware.Logger())
	}
	// e.OnAddRouteHandler = func(host string, route echo.Route, handler echo.HandlerFunc, middleware []echo.MiddlewareFunc) {
	// 	fmt.Println(host, route.Path)
	// }

	return &Route{
		Echo:    e,
		log:     dh.Logger,
		service: service,
	}
}

func (route *Route) RegisterRoutes() {
	route.AuthRoutes()
	route.UserGroupRoute()
}

func (route *Route) Start(port string) error {
	if err := route.Echo.Start(port); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (route *Route) Shutdown(ctx context.Context) error {
	if err := route.Echo.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
