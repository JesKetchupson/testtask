// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type NewPost struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Post struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	PublicationDate time.Time `json:"publicationDate"`
}
