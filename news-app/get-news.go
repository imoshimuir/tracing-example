package main

import (
	"context"
	"latest-news/types"
)

type GetNews func(ctx context.Context) ([]types.News, error)

func NewGetNews(store Store) GetNews {
	return func(ctx context.Context) ([]types.News, error) {
		results, err := store.Get(ctx)
		if err != nil {
			return nil, err
		}
		return results, nil
	}
}
