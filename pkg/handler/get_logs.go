package handler

import (
	"imserver/pkg/auth"
	"imserver/pkg/store"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetTotal(logStore store.LogStoreType) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !auth.CanRead(c.Get("roles").(string)) {
			return echo.ErrForbidden
		}

		total, err := strconv.Atoi(c.QueryParam("filter"))
		if err != nil {
			total = -1
		}

		rows, err := logStore.GetLogs(auth.GetSourceID(c), total)
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, rows)
	}
}
