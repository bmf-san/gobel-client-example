package model

// A Comments represents the plural of comment.
type Comments []Comment

// A Comment represents the singular of comment.
type Comment struct {
	ID     int    `json:"id"`
	PostID int    `json:"post"`
	Body   string `json:"body"`
}
