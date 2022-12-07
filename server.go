package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "imserver/docs"
	"imserver/pkg/db"
	"imserver/pkg/handler"
	imserverMiddleware "imserver/pkg/middleware"
	"imserver/pkg/store"
	"imserver/pkg/validator"
	"os"
)

// @title Swagger Example API
// @version 1.0
// @description This is a REST API for storing logs in immudb from different services.

// @contact.name mirpo
// @contact.url https://github.com/mirpo

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 0.0.0.0
// @BasePath /v1
func main() {
	e := echo.New()
	e.Validator = validator.NewValidator()
	e.Debug = true

	dbClient := db.NewClient()
	defer dbClient.CloseSession(context.TODO())

	// init db access
	logStore := store.NewLogStore(dbClient)
	sourceStore := store.NewSourceStore(dbClient)

	// Middleware
	e.Use(middleware.Recover())
	e.Use(imserverMiddleware.LoggerConfig())
	e.Use(imserverMiddleware.CorsConfig())
	e.Use(imserverMiddleware.JWTConfig(sourceStore))

	// swagger documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Routes in default v1 group
	g := e.Group("/v1")
	g.GET("/health", handler.HealthCheck)
	g.POST("/logs", handler.CreateLog(logStore))        // Store single log line
	g.POST("/logs/batch", handler.CreateLogs(logStore)) // Store batch of log lines
	g.GET("/logs", handler.GetTotal(logStore))          // Print history of stored logs (all, last x)
	g.GET("/logs/count", handler.GetCount(logStore))    // Print number of stored logs

	// Start server
	httpPort := os.Getenv("IMSERVER_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	e.Logger.Fatal(e.Start(":" + httpPort))
}
