package http

import (
	"github.com/dagulv/train-api/internal/adapter/http/routes"
	"github.com/labstack/echo/v4"
)

func (s Server) addRoutes(e *echo.Echo) {
	routes.Routes(e, s.User, s.Json)
}
