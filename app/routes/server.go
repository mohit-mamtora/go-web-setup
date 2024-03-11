package routes

import (
	"context"
	"net/http"

	"github.com/Mohit-Mamtora/gofinlop/app"
	"github.com/Mohit-Mamtora/gofinlop/app/logger"
	"github.com/Mohit-Mamtora/gofinlop/app/services"
	"github.com/labstack/echo/v4"
)

type (
	Route struct {
		*echo.Echo
		log     logger.Log
		service *services.Service
	}
)

func InitilizeRoute(service *services.Service, dh *app.DependencyHandler) *Route {
	e := echo.New()

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
