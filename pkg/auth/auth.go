package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"strings"
)

type Claims struct {
	SourceID int64 `json:"sourceId"`
	jwt.StandardClaims
}

func SkipperFn(c echo.Context) bool {
	path := c.Request().URL.Path
	return path == "/v1/health" || strings.Contains(path, "swagger")
}

func GetSourceID(c echo.Context) int64 {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*Claims)
	return claims.SourceID
}

func CanWrite(roles string) bool {
	return strings.Contains(roles, "ROLE_WRITE")
}

func CanRead(roles string) bool {
	return strings.Contains(roles, "ROLE_READ")
}
