package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func CorsConfig() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("IMSERVER_CORS")},
		AllowMethods: []string{echo.GET, echo.POST},
	})
}
