package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {
	e := echo.New()

	t.Run("SkipperFn", func(t *testing.T) {
		t.Run("failure", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			context := e.NewContext(req, rec)

			require.False(t, SkipperFn(context))
		})

		t.Run("success", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/v1/health", nil)
			rec := httptest.NewRecorder()
			context := e.NewContext(req, rec)

			require.True(t, SkipperFn(context))
		})
	})
}
