package handler

import (
	"github.com/labstack/echo/v4"
	"imserver/pkg/auth"
	"imserver/pkg/store"
	"net/http"
)

type TotalResponse struct {
	Count int64 `json:"count"`
}

func GetCount(logStore store.LogStoreType) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !auth.CanRead(c.Get("roles").(string)) {
			return echo.ErrForbidden
		}

		rows, err := logStore.GetTotal(auth.GetSourceID(c))
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSON(http.StatusOK, rows)
	}
}
