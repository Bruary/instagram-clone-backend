package models

import "time"

type Comment struct {
	Id             int
	Uid            string
	Comment_string string
	Likes_count    int
	User_id        int
	Parent_id      int
	Post_id        int
	Created_at     time.Time
	Updated_at     time.Time
	Deleted_at     time.Time
}
