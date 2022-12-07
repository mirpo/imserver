package handler

import (
	"imserver/pkg/auth"
	"imserver/pkg/model"
	"imserver/pkg/store"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type LogRequest struct {
	Metrics string `json:"metrics" validate:"required"`
}

func CreateLog(logStore store.LogStoreType) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !auth.CanWrite(c.Get("roles").(string)) {
			return echo.ErrForbidden
		}

		logReq := LogRequest{}

		if err := c.Bind(&logReq); err != nil {
			c.Logger().Error(err)
			return echo.ErrBadRequest
		}
		if err := c.Validate(&logReq); err != nil {
			c.Logger().Error(err)
			return err
		}

		log := model.Log{
			SourceID:  auth.GetSourceID(c),
			CreatedAT: time.Now().Unix(),
			Metrics:   logReq.Metrics,
		}

		if err := logStore.CreateLog(log); err != nil {
			c.Logger().Error(err)
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusCreated, log)
	}
}
