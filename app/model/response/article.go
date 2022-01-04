package response

import "time"

type Comment struct {
	Id        interface{} `json:"id,i"`
	Author    string      `json:"author,omitempty"`
	Comment   string      `json:"comment,omitempty"`
	CreatedAt time.Time   `json:"created_at"`
}

type Article struct {
	Id        string    `json:"id" bson:"_id"`
	Author    string    `json:"author,omitempty"`
	Title     string    `json:"title,omitempty"`
	Body      string    `json:"body,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Comments  []Comment `json:"comments,omitempty"`
}
