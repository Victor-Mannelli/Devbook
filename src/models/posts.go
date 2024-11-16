package models

import "time"

type Post struct {
	PostId         uint64    `json:"post_id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorId       uint64    `json:"author_id,omitempty"`
	AuthorUsername string    `json:"author_name,omitempty"`
	Likes          uint64    `json:"likes"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}
