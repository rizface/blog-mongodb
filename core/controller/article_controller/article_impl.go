package article_controller

import (
	"blog-mongo/app/exception"
	request2 "blog-mongo/app/model/request"
	"blog-mongo/core/controller"
	"blog-mongo/core/service/article_service"
	"blog-mongo/helper"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type article struct {
	service article_service.Article
}

func NewArticle(service article_service.Article) controller.BasicCrud {
	return &article{service: service}
}

func (a article) Get(w http.ResponseWriter, r *http.Request) {
	var page int
	var err error
	pageParam := helper.GetQueryValue(r,"page")
	if len(pageParam) == 0 || pageParam == "0"{
		page = 1
	} else {
		page,err = strconv.Atoi(pageParam)
		helper.PanicCustomException(exception.BadRequest{Err: errors.New("halaman tidak ditemukan")},err != nil)
	}
	result := a.service.Get(int64(page))
	helper.JsonWriter(w,http.StatusOK,"success",map[string]interface{} {
		"next": "?page="+strconv.Itoa(page + 1),
		"prev": "?page="+strconv.Itoa(page - 1),
		"articles": result,
	})
}

func (a article) Post(w http.ResponseWriter, r *http.Request) {
	var request request2.Article
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Panic(err)
	author := helper.GetClaimsValue(r,"userData","username")
	request.Author = author.(string)
	result := a.service.Post(request)
	helper.JsonWriter(w,http.StatusOK,result,nil)
}

func (a article) Delete(w http.ResponseWriter, r *http.Request) {
	articleId := helper.GetParamsValue(r,"articleId")
	author := helper.GetClaimsValue(r,"userData","username")
	result := a.service.Delete(author.(string),articleId)
	helper.JsonWriter(w,http.StatusOK,result,nil)
}

func (a article) Update(w http.ResponseWriter, r *http.Request) {
	var request request2.UpdateArticle
	articleId := helper.GetParamsValue(r,"articleId")
	author := helper.GetClaimsValue(r,"userData","username")
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Panic(err)
	request.Id,request.Author = articleId,author.(string)
	result := a.service.Update(request)
	helper.JsonWriter(w,http.StatusOK,result,nil)
}



