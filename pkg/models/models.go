package models

type User struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Rating int64  `json:"rating"`
}
