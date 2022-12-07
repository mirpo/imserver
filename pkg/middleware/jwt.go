package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"imserver/pkg/auth"
	"imserver/pkg/store"
	"os"
)

func JWTConfig(sourceStore store.SourceStoreType) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper:    auth.SkipperFn,
		Claims:     &auth.Claims{},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		SuccessHandler: func(c echo.Context) {
			c.Set("roles", "") // by default deny *
			source, err := sourceStore.Get(auth.GetSourceID(c))
			if err == nil {
				c.Set("roles", source.Roles) // by default deny *
			}
		},
	})
}
