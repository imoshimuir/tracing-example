package main

import (
	"context"
	"latest-news/types"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type Store interface {
	Get(ctx context.Context) ([]types.News, error)
}

// TracedStore is a tracing wrapper implementation for Store.
type TracedStore struct {
	tracer trace.Tracer
	inner  Store
}

// Get traces a call to the same method on t.inner.
func (t *TracedStore) Get(ctx context.Context) ([]types.News, error) {
	ctx, span := t.tracer.Start(ctx, "Postgres.Get")
	defer span.End()
	list, err := t.inner.Get(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
	span.SetAttributes(attribute.Int("result.len", len(list)))

	return list, err
}

// NewTracedStore wraps calls to inner with tracing spans using
// tracer.
func NewTracedStore(tracer trace.Tracer, inner Store) *TracedStore {
	return &TracedStore{
		tracer: tracer,
		inner:  inner,
	}
}
