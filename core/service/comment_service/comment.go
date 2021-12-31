package comment_service

import "blog-mongo/app/model/request"

type Comment interface {
	Post(request request.Comment) string
	Delete(request request.DeleteComment) string
	Update(request request.UpdateComment) string
}
