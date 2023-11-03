package db

import (
	"context"
	"database/sql"
	"latest-news/types"

	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func New(ctx context.Context, connectionString string) (*Store, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return &Store{}, err
	}

	if err := db.Ping(); err != nil {
		return &Store{}, err
	}

	return &Store{db: db}, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) Get(ctx context.Context) ([]types.News, error) {
	rows, err := s.db.Query("SELECT title, description, timestamp, id FROM news ORDER BY timestamp DESC LIMIT 5")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var newsList []types.News

	for rows.Next() {
		var news types.News
		if err := rows.Scan(&news.Title, &news.Description, &news.Timestamp, &news.ID); err != nil {
			return nil, err
		}
		newsList = append(newsList, news)
	}

	return newsList, nil
}
