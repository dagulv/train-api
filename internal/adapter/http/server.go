package http

import (
	"context"
	"net/http"

	"github.com/dagulv/train-api/internal/domain/user"
	"github.com/go-webauthn/webauthn/webauthn"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Json     jsoniter.API
	WebAuthn *webauthn.WebAuthn
	User     user.Service
}

func (s Server) StartServer(ctx context.Context) (err error) {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	s.addRoutes(e)
	return e.Start(":3001")
}
