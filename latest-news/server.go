package main

import (
	"context"
	"fmt"
	"latest-news/db"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/trace"
)

func InstantiateServer(ctx context.Context,
	cfg *GlobalConfig, tracer trace.Tracer) (*echo.Echo, error) {
	store, err := db.New(ctx, cfg.PostgresConnectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate store: %w", err)
	}

	e := echo.New()

	// Adds the OpenTelemetry middleware to all routes
	e.Use(otelecho.Middleware(serviceName))

	getNews := NewGetNews(NewTracedStore(tracer, store))

	e.GET("/news", GetNewsHandler(ctx, getNews))
	return e, nil
}

func GetNewsHandler(ctx context.Context, GetNews GetNews) echo.HandlerFunc {
	return func(c echo.Context) error {
		news, err := GetNews(c.Request().Context())
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to retrieve news")
		}

		if len(news) == 0 {
			return c.String(http.StatusNotFound, "No news available")
		}

		return c.JSON(http.StatusOK, news)
	}
}
