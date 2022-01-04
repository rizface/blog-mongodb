package comment_service

import "blog-mongo/app/model/request"

type Comment interface {
	GetById(articleId string, commentId string) interface{}
	Post(request request.Comment) string
	Delete(request request.DeleteComment) string
	Update(request request.UpdateComment) string
}
