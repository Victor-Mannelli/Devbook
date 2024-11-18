package models

import "time"

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	Followers []User    `json:"followers"`
	Followed  []User    `json:"followed"`
	Posts     []Post    `json:"posts"`
}
