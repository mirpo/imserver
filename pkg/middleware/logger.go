package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggerConfig() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]    ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	})
}
