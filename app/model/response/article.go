package response

import "time"

type Comment struct {
	Id        interface{} `json:"id"`
	Author    string      `json:"author"`
	Comment   string      `json:"comment"`
	CreatedAt time.Time   `json:"created_at"`
}

type Article struct {
	Id        string    `json:"id" bson:"_id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	Comments  []Comment `json:"comments"`
}
