package types

type NewsFeed []News

type News struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Timestamp   int64  `json:"timestamp"`
	ID          string `json:"id"`
}
