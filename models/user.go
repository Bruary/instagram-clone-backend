package models

import "time"

type NewUser struct {
	Name     string
	Email    string
	Password string
}

type User struct {
	Id                int
	Name              string
	Email             string
	Password          string
	Profile_image_url string
	Profile_public_id string
	Active            bool
	Verified_at       time.Time
	Created_at        time.Time
	Updated_at        time.Time
	Deleted_at        time.Time
}
