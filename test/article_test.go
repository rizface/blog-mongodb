package test

import (
	"blog-mongo/app/model/request"
	"blog-mongo/app/setup"
	"blog-mongo/core/repository/article_repository"
	"blog-mongo/core/service/article_service"
	"blog-mongo/helper"
	"blog-mongo/route"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestPostArticle(t *testing.T) {
	token := helper.TestingToken()
	t.Run("repo", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			db,_ := helper.Connection()
			ctx,cancel := helper.CreateCtx(10)
			defer cancel()
			repo := article_repository.NewArticle()
			result,err := repo.Post(db,ctx,request.Article{
				Author:    "fariz",
				Title:     "METAVERSE",
				Body:      "facebook bikin metarverse",
				CreatedAt: time.Now(),
			})
			assert.Nil(t, err)
			assert.NotNil(t, result)
		})
	})
	t.Run("service", func(t *testing.T) {
		db,_ := helper.Connection()
		valid := validator.New()
		repo := article_repository.NewArticle()
		service := article_service.NewArticle(valid,db,repo)
		t.Run("success", func(t *testing.T) {
			result := service.Post(request.Article{
				Author:    "fariz",
				Title:     "Indonesia vs Thailand",
				Body:      "indonesia lawan thailand",
				CreatedAt: time.Now(),
			})
			assert.True(t, len(result) > 0)
		})
	})
	t.Run("controller", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			router := setup.ArticleRouter()
			payload := request.Article{
				Title:     "JWT",
				Body:      "Json Web Token",
			}
			payloadJson,_ := json.Marshal(payload)
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPost,route.ARTIKEL,bytes.NewReader(payloadJson))
			request.Header.Add("Authorization", "Bearer " + token)
			router.ServeHTTP(recorder,request)
			assert.Equal(t, recorder.Code,http.StatusOK)
		})
		t.Run("failed_empty_request", func(t *testing.T) {
			router := setup.ArticleRouter()
			payload := request.Article{}
			payloadJson,_ := json.Marshal(payload)
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPost,route.ARTIKEL,bytes.NewReader(payloadJson))
			request.Header.Add("Authorization", "Bearer " + token)
			router.ServeHTTP(recorder,request)
			result,_ := ioutil.ReadAll(recorder.Body)
			fmt.Println(string(result))
			assert.Equal(t, recorder.Code,http.StatusBadRequest)
		})
	})
}

func TestGetArticle(t *testing.T)  {
	repo := article_repository.NewArticle()
	valid := validator.New()
	db,_ := helper.Connection()
	t.Run("service", func(t *testing.T) {
		service := article_service.NewArticle(valid,db,repo)
		service.Get(1)
	})
}

func TestDeleteArticle(t *testing.T) {
	router := setup.ArticleRouter()
	token := helper.TestingToken()
	t.Run("success", func(t *testing.T) {
		id := helper.PostArticleTesting()
		recorder := httptest.NewRecorder()
		request :=  httptest.NewRequest(http.MethodDelete,"/articles/" + id,nil)
		request.Header.Add("Authorization", "Bearer " + token)
		router.ServeHTTP(recorder,request)
		assert.Equal(t, http.StatusOK,recorder.Code)
	})
	t.Run("failed", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request :=  httptest.NewRequest(http.MethodDelete,"/articles/61cc09e33e02de1ce141ec00",nil)
		request.Header.Add("Authorization", "Bearer " + token)
		router.ServeHTTP(recorder,request)
		assert.Equal(t, http.StatusInternalServerError,recorder.Code)
	})
}

func TestUpdateArticle(t *testing.T) {
	router := setup.ArticleRouter()
	token := helper.TestingToken()
	t.Run("success", func(t *testing.T) {
		id := helper.PostArticleTesting()
		payload := request.UpdateArticle{
			Title:  "Pembaharuan Energy",
			Body:   "Energy Terbarukan",
		}
		payloadJson,_ := json.Marshal(payload)
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPut,"/articles/" + id,bytes.NewReader(payloadJson))
		request.Header.Add("Authorization","Bearer " + token)
		router.ServeHTTP(recorder,request)
		assert.Equal(t, http.StatusOK,recorder.Code)
	})
	t.Run("failed", func(t *testing.T) {
		payload := request.UpdateArticle{
			Title:  "Pembaharuan Energy",
			Body:   "Energy Terbarukan",
		}
		payloadJson,_ := json.Marshal(payload)
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPut,"/articles/61cc09e33e02de1ce141ec00",bytes.NewReader(payloadJson))
		request.Header.Add("Authorization","Bearer " + token)
		router.ServeHTTP(recorder,request)
		assert.Equal(t, http.StatusInternalServerError,recorder.Code)
	})
}