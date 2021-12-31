package test

import (
	request2 "blog-mongo/app/model/request"
	"blog-mongo/app/setup"
	"blog-mongo/helper"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var token = helper.TestingToken()
var router = setup.CommentRouter()


func TestPostComment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		articleId,_ := helper.PostTestingComment()
		payload := request2.Comment{
			Comment:   "INI KOMENTAR PERTAMA",
		}
		payloadJson,_ := json.Marshal(payload)
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost,fmt.Sprintf("/articles/%s/comments",articleId),bytes.NewReader(payloadJson))
		request.Header.Add("Authorization", "Bearer " + token)
		router.ServeHTTP(recorder,request)
		assert.Equal(t, http.StatusOK,recorder.Code)
	})
	t.Run("not found", func(t *testing.T) {
		payload := request2.Comment{
			Comment:   "INI KOMENTAR PERTAMA",
		}
		payloadJson,_ := json.Marshal(payload)
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost,"/articles/61cd8a2ca4f1680c2ed36cb9/comments",bytes.NewReader(payloadJson))
		request.Header.Add("Authorization", "Bearer " + token)
		router.ServeHTTP(recorder,request)
		assert.Equal(t, http.StatusNotFound,recorder.Code)
	})
	t.Run("bad request", func(t *testing.T) {
		payload := request2.Comment{}
		payloadJson,_ := json.Marshal(payload)
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost,"/articles/61cd8a2ca4f1680c2ed36cb9/comments",bytes.NewReader(payloadJson))
		request.Header.Add("Authorization", "Bearer " + token)
		router.ServeHTTP(recorder,request)
		assert.Equal(t, http.StatusBadRequest,recorder.Code)
	})
}

func TestDeleteComment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		articleId,commentId := helper.PostTestingComment()
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodDelete,fmt.Sprintf("/articles/%s/comments/%s",articleId,commentId),nil)
		request.Header.Add("Authorization", "Bearer " + token)
		router.ServeHTTP(recorder,request)
		assert.Equal(t, http.StatusOK,recorder.Code)
	})
	t.Run("not found artikel", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodDelete,"/articles/61cd8a2ca4f1680c2ed36cb2/comments/61cd9c7b156288f11f6c36da",nil)
		request.Header.Add("Authorization", "Bearer " + token)
		router.ServeHTTP(recorder,request)
		assert.Equal(t, http.StatusNotFound,recorder.Code)
	})
	t.Run("not found comment", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodDelete,"/articles/61cd8a2ca4f1680c2ed36cb3/comments/61cd9c7b156288f11f6c36db",nil)
		request.Header.Add("Authorization", "Bearer " + token)
		router.ServeHTTP(recorder,request)
		assert.Equal(t, http.StatusNotFound,recorder.Code)
	})
}

func TestUpdateComment(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		articleId,commentId := helper.PostTestingComment()
		payload := request2.UpdateComment{
			ArticleId: articleId,
			CommentId: commentId,
			Comment:   "KOMENTAR DARI TESTING NI BOSS",
		}
		payloadJson,_ := json.Marshal(payload)
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPut,fmt.Sprintf("/articles/%s/comments/%s",articleId,commentId),bytes.NewReader(payloadJson))
		request.Header.Add("Authorization","Bearer " + token)
		router.ServeHTTP(recorder,request)
		assert.Equal(t, http.StatusOK,recorder.Code)
	})
}