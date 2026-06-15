package main

import (
	"time"
)

type Priority int

const (
	Low Priority = iota
	Medium
	High
)

type Snippet struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	Priority  Priority  `json:"priority"`
}

func main() {

}
