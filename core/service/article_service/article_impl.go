package article_service

import (
	"blog-mongo/app/exception"
	"blog-mongo/app/model/request"
	"blog-mongo/app/model/response"
	"blog-mongo/core/repository/article_repository"
	"blog-mongo/helper"
	"errors"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type article struct {
	valid *validator.Validate
	db    *mongo.Database
	repo  article_repository.Article
}

func NewArticle(valid *validator.Validate, db *mongo.Database, repo article_repository.Article) Article {
	return &article{
		valid: valid,
		db:    db,
		repo:  repo,
	}
}

func (a *article) Get(page int64) []response.Article {
	var result []response.Article
	var skip int64
	var limit int64 = 15
	ctx,cancel := helper.CreateCtx(10)
	defer cancel()
	skip = (page - 1) * 15
	cursor,err := a.repo.Get(a.db,ctx,skip,limit)
	helper.Panic(err)
	cCtx,cCancel := helper.CreateCtx(5)
	defer cCancel()
	err = cursor.All(cCtx,&result)
	helper.Panic(err)
	return result
}

func (a *article) GetById(articleId string) response.Article {
	var result response.Article
	valid := primitive.IsValidObjectID(articleId)
	helper.PanicCustomException(exception.NotFound{Err: errors.New("artikel tidak ditemukan")},valid == false)
	ctx,cancel := helper.CreateCtx(5)
	defer cancel()
	cursor := a.repo.GetById(a.db,ctx,articleId)
	err := cursor.Decode(&result)
	helper.PanicCustomException(exception.NotFound{Err:errors.New("artikel tidak ditemukan")},err != nil && errors.Is(err,mongo.ErrNoDocuments))
	helper.PanicCustomException(exception.InternalError{Err:errors.New("terjadi kesalahan pada sistem kami")},err != nil)
	return result
}

func (a *article) Post(article request.Article) string {
	err := a.valid.Struct(article)
	helper.PanicCustomException(exception.BadRequest{Err: errors.New("Artikel Tidak Valid")},err != nil)
	ctx,cancel := helper.CreateCtx(10)
	defer cancel()
	_,err = a.repo.Post(a.db,ctx,article)
	helper.PanicCustomException(exception.InternalError{Err:errors.New("terjadi kesalahan pada sistem kami")},err != nil)
	return "artikel berhasil ditambahkan"
}

func (a *article) Delete(author string,articleId string) string {
	ctx,cancel := helper.CreateCtx(10)
	defer cancel()
	valid := primitive.IsValidObjectID(articleId)
	helper.PanicCustomException(exception.NotFound{Err:errors.New("artikel tidak ditemukan")}, valid == false)
	success,err := a.repo.Delete(a.db,ctx,author,articleId)
	helper.PanicCustomException(exception.InternalError{Err:errors.New("terjadi kesalahan pada sistem kami")},success == false || err != nil)
	return "artikel berhasil dihapus"
}

func (a *article) Update(request request.UpdateArticle) string {
	err := a.valid.Struct(request)
	helper.PanicCustomException(exception.BadRequest{Err:errors.New("artikel tidak valid")},err != nil)
	valid := primitive.IsValidObjectID(request.Id)
	helper.PanicCustomException(exception.NotFound{Err:errors.New("artikel tidak ditemukan")}, valid == false)
	ctx,cancel := helper.CreateCtx(10)
	defer cancel()
	success,err := a.repo.Update(a.db,ctx,request)
	helper.PanicCustomException(exception.InternalError{Err:errors.New("terjadi kesalahan pada sistem kami")},success == false || err != nil)
	return "artikel berhasil diupdate"
}
