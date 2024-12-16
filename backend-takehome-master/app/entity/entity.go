package entity

import "time"

type User struct {
	Name         string
	Email        string
	HashPassword string
	Password     string
	UserID       int
}

type BlogPost struct {
	Title       string
	Content     string
	AuthorID    int
	AuthorName  string
	PostID      int
	CreatedTime time.Time
	UpdatedTime time.Time
	Comments    []Comment
}

type Comment struct {
	CommentID   int
	PostID      int
	AuthorID    int
	Author      string
	Content     string
	CreatedTime time.Time
}
