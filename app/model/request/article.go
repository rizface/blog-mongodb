package request

import "time"

type Article struct {
	Author    string    `json:"author" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	Body      string    `json:"body" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
}

type UpdateArticle struct {
	Id     string `validate:"required"`
	Author string `validate:"required"`
	Title  string `json:"title" validate:"required"`
	Body   string `json:"body" validate:"required"`
}
