package todo

import "time"

type ID string

type Todo struct {
	Id        ID        `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}
