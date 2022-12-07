package handler

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"imserver/pkg/auth"
	storeMock "imserver/pkg/store/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTotal(t *testing.T) {
	e := echo.New()
	logStore := &storeMock.LogStoreMock{}
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	t.Run("error when failed to GetLogs", func(t *testing.T) {
		context := e.NewContext(req, rec)
		context.Set("user", &jwt.Token{
			Claims: &auth.Claims{
				SourceID:       12345,
				StandardClaims: jwt.StandardClaims{},
			},
		})
		context.Set("roles", "")
		handler := GetTotal(logStore)

		res := handler(context)

		require.Error(t, res)
	})

	t.Run("success", func(t *testing.T) {
		context := e.NewContext(req, rec)
		context.Set("user", &jwt.Token{
			Claims: &auth.Claims{
				SourceID:       1,
				StandardClaims: jwt.StandardClaims{},
			},
		})
		context.Set("roles", "ROLE_READ")
		handler := GetTotal(logStore)

		res := handler(context)

		require.Nil(t, res)
		require.Equal(t, "[{\"source_id\":1,\"created_at\":12345,\"metrics\":\"metric1\"},{\"source_id\":1,\"created_at\":23456,\"metrics\":\"metric2\"}]\n", rec.Body.String())
	})
}
