package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Msg string `json:"msg"`
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags health
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{Msg: "Service is healthy!"})
}
