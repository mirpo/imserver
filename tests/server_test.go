package tests

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"imserver/pkg/model"
	"net/http"
	"testing"
)

// read & write
const source1Jwt = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzb3VyY2VJZCI6MX0.5VtXO9J1YF2sv8SwTfvsVseqHMjEwhFBHJLpSuj-i34"

// only write
const source2Jwt = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzb3VyY2VJZCI6Mn0.KxzbtbC6E8TNt0NmmRBdNz5P9ixj6sSKE9JQVk3fkGg"

// no permissions
const source3Jwt = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzb3VyY2VJZCI6M30.spUBA7KeBfZqevVxwfKGUAxXFHOJGvpOxV6-x339d-M"

func TestServer(t *testing.T) {
	client := resty.New().
		SetBaseURL("http://0.0.0.0:1323/v1").
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json")

	t.Run("Health endpoint", func(t *testing.T) {
		t.Run("200", func(t *testing.T) {
			resp, _ := client.R().
				Get("/health")

			require.Equal(t, http.StatusOK, resp.StatusCode())
			require.Equal(t, "{\n  \"msg\": \"Service is healthy!\"\n}\n", string(resp.Body()))
		})
	})

	t.Run("Create log", func(t *testing.T) {
		t.Run("400 Bad Request", func(t *testing.T) {
			resp, _ := client.R().
				SetBody(`{"metrics":"xyz"}`).
				Post("/logs")

			require.Equal(t, http.StatusBadRequest, resp.StatusCode())
		})

		t.Run("401 Unauthorized", func(t *testing.T) {
			resp, _ := client.R().
				SetBody(`{"metrics":"xyz"}`).
				SetAuthToken("TOKEN").
				Post("/logs")

			require.Equal(t, http.StatusUnauthorized, resp.StatusCode())
		})

		t.Run("201 Created", func(t *testing.T) {
			resp, _ := client.R().
				SetBody(`{"metrics":"xyz"}`).
				SetAuthToken(source1Jwt).
				Post("/logs")

			require.Equal(t, http.StatusCreated, resp.StatusCode())

			var log model.Log
			_ = json.Unmarshal(resp.Body(), &log)
			require.Equal(t, int64(1), log.SourceID)
			require.Equal(t, "xyz", log.Metrics)
		})
	})

	t.Run("Roles", func(t *testing.T) {
		t.Run("403 if no roles are specified", func(t *testing.T) {
			resp, _ := client.R().
				SetBody(`{"metrics":"xyz"}`).
				SetAuthToken(source3Jwt).
				Post("/logs")

			require.Equal(t, http.StatusForbidden, resp.StatusCode())
		})

		t.Run("403 if read role is not set", func(t *testing.T) {
			resp, _ := client.R().
				SetBody(`{"metrics":"xyz"}`).
				SetAuthToken(source2Jwt).
				Get("/logs")

			require.Equal(t, http.StatusForbidden, resp.StatusCode())
		})

		t.Run("but 201 because write roles is set", func(t *testing.T) {
			resp, _ := client.R().
				SetBody(`{"metrics":"xyz"}`).
				SetAuthToken(source2Jwt).
				Post("/logs")

			require.Equal(t, http.StatusCreated, resp.StatusCode())
		})
	})
}
