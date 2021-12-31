package request

import "time"

type Comment struct {
	ArticleId string `validate:"required"`
	Author    string `validate:"required"`
	Comment   string `validate:"required"`
	CreatedAt time.Time
}

type DeleteComment struct {
	ArticleId string `validate:"required"`
	CommentId string `validate:"required"`
	Username  string `validate:"required"`
}

type UpdateComment struct {
	ArticleId string `validate:"required"`
	CommentId string `validate:"required"`
	Username  string `validate:"required"`
	Comment   string `json:"comment" validate:"required"`
}
