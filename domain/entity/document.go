package entity

import "time"

type Document struct {
	Id          int       `json:"id"`
	Uid         string    `json:"uid"`
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	DateCreated time.Time `json:"date_created"`
	Creator     User      `json:"creator"`
	Status      Status    `json:"status"`
	Project     Project   `json:"project"`
	Space       Space     `json:"space"`
}
