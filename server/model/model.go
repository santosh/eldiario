package model

import "time"

// Entry Struct (Model)
type Entry struct {
	// TODO: Can we please use mongo's native _id?
	ID       string    `json:"id"`
	Body     string    `json:"body"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Author   string    `json:"author"`
}
