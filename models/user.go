package models

import "time"

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id                int
	UUID              string
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
