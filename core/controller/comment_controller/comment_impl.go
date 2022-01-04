package comment_controller

import (
	"blog-mongo/app/exception"
	request2 "blog-mongo/app/model/request"
	"blog-mongo/core/controller"
	"blog-mongo/core/service/comment_service"
	"blog-mongo/helper"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type comment struct {
	service comment_service.Comment
}

func NewComment(service comment_service.Comment) controller.BasicCrud {
	return &comment{service: service}
}

func (c *comment) Get(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (c *comment) GetById(w http.ResponseWriter, r *http.Request) {
	articleId := helper.GetParamsValue(r,"articleId")
	commentId := helper.GetParamsValue(r,"commentId")
	result := c.service.GetById(articleId,commentId)
	helper.JsonWriter(w,http.StatusOK,"success",map[string]interface{} {
		"articleId":articleId,
		"commentId":commentId,
		"comment": result,
	})
}

func (c *comment) Post(w http.ResponseWriter, r *http.Request) {
	var request request2.Comment
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Panic(err)
	request.Author = helper.GetClaimsValue(r,"userData","username").(string)
	request.ArticleId = helper.GetParamsValue(r,"articleId")
	request.CreatedAt = time.Now()
	result := c.service.Post(request)
	helper.JsonWriter(w,http.StatusOK,result,nil)
}

func (c *comment) Delete(w http.ResponseWriter, r *http.Request) {
	var request request2.DeleteComment
	request.Username = helper.GetClaimsValue(r,"userData","username").(string)
	request.CommentId = helper.GetParamsValue(r,"commentId")
	request.ArticleId = helper.GetParamsValue(r,"articleId")
	result := c.service.Delete(request)
	helper.JsonWriter(w,http.StatusOK,result,nil)
}

func (c *comment) Update(w http.ResponseWriter, r *http.Request) {
	var request request2.UpdateComment
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.PanicCustomException(exception.InternalError{Err:errors.New("terjadi kesalahan pada sistem kamu")},err != nil)
	request.Username = helper.GetClaimsValue(r,"userData","username").(string)
	request.CommentId = helper.GetParamsValue(r,"commentId")
	request.ArticleId = helper.GetParamsValue(r,"articleId")
	result := c.service.Update(request)
	helper.JsonWriter(w,http.StatusOK,result,nil)
}

