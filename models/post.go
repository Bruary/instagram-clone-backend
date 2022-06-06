package models

import "time"

type Post struct {
	Id          int
	UID         string
	Post_url    string
	Caption     string
	Likes_count int
	User_id     int
	Created_at  time.Time
	Updated_at  time.Time
	Deleted_at  time.Time
}
