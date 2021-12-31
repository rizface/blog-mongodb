package comment_service

import (
	"blog-mongo/app/exception"
	"blog-mongo/app/model/request"
	"blog-mongo/core/repository/comment_respository"
	"blog-mongo/helper"
	"errors"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type comment struct {
	db    *mongo.Database
	valid *validator.Validate
	repo  comment_respository.Comment
}

func NewComment(db *mongo.Database, valid *validator.Validate, repo comment_respository.Comment) Comment {
	return &comment{
		db:    db,
		valid: valid,
		repo:  repo,
	}
}

func (c *comment) Post(request request.Comment) string {
	err := c.valid.Struct(request)
	helper.PanicCustomException(exception.BadRequest{Err:errors.New("komentar tidak valid")},err != nil)
	valid := primitive.IsValidObjectID(request.ArticleId)
	helper.PanicCustomException(exception.NotFound{Err:errors.New("artikel tidak ditemukan")}, valid == false)
	ctx,cancel := helper.CreateCtx(5)
	defer cancel()
	_,success,err := c.repo.Post(c.db,ctx,request)
	//fmt.Println("SUCCESS :", success,"ERR : ",err)
	//time.Sleep(10 * time.Second)
	helper.PanicCustomException(exception.NotFound{Err:errors.New("artikel tidak ditemukan")}, success == false)
	helper.PanicCustomException(exception.InternalError{Err:errors.New("terjadi kesalahan pada sistem kami")}, err != nil)
	return "komentar berhasil diposting"
}

func (c *comment) Delete(request request.DeleteComment) string {
	err := c.valid.Struct(request)

	helper.PanicCustomException(exception.BadRequest{Err:errors.New("input tidak valid")},err != nil)
	valid := primitive.IsValidObjectID(request.ArticleId)
	helper.PanicCustomException(exception.NotFound{Err:errors.New("article tidak ditemukan")}, valid == false)

	valid = primitive.IsValidObjectID(request.ArticleId)
	helper.PanicCustomException(exception.NotFound{Err:errors.New("komentar tidak ditemukan")}, valid == false)

	ctx,cancel := helper.CreateCtx(10)
	defer cancel()
	success,err := c.repo.Delete(c.db,ctx,request)
	helper.PanicCustomException(exception.NotFound{Err:errors.New("terjadi kesalahan pada sistem kami")},success == false || err != nil)
	return "komentar berhasil dihapus"
}

func (c *comment) Update(request request.UpdateComment) string {
	err := c.valid.Struct(request)

	helper.PanicCustomException(exception.BadRequest{Err:errors.New("input tidak valid")},err != nil)
	valid := primitive.IsValidObjectID(request.ArticleId)
	helper.PanicCustomException(exception.NotFound{Err:errors.New("article tidak ditemukan")}, valid == false)

	valid = primitive.IsValidObjectID(request.ArticleId)
	helper.PanicCustomException(exception.NotFound{Err:errors.New("komentar tidak ditemukan")}, valid == false)

	ctx,cancel := helper.CreateCtx(10)
	defer cancel()
	success,err := c.repo.Update(c.db,ctx,request)
	helper.Panic(err)
	helper.PanicCustomException(exception.InternalError{Err:errors.New("terjadi kesalahan pada sistem kami")},success == false)
	return "komentar berhasil diupdate"
}

