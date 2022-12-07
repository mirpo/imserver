package handler

import (
	"imserver/pkg/auth"
	"imserver/pkg/model"
	"imserver/pkg/store"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type LogsBatchRequest []LogRequest

func CreateLogs(logStore store.LogStoreType) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !auth.CanWrite(c.Get("roles").(string)) {
			return echo.ErrForbidden
		}

		batchReq := LogsBatchRequest{}
		c.Logger().Debugf("sourceId: %v", auth.GetSourceID(c))
		if err := c.Bind(&batchReq); err != nil {
			return echo.ErrBadRequest
		}

		// validate each log
		for i := range batchReq {
			if err := c.Validate(&batchReq[i]); err != nil {
				return err
			}
		}

		c.Logger().Debugf("batch logs request: %v", batchReq)
		c.Logger().Debugf("sourceId: %v", auth.GetSourceID(c))

		var logs []model.Log

		for _, req := range batchReq {
			log := model.Log{
				SourceID:  auth.GetSourceID(c),
				CreatedAT: time.Now().Unix(),
				Metrics:   req.Metrics,
			}
			logs = append(logs, log)
		}

		if err := logStore.CreateLogs(logs); err != nil {
			c.Logger().Error(err)
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusCreated, logs)
	}
}
