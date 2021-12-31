package setup

import (
	"blog-mongo/app/exception"
	"blog-mongo/core/controller"
	"blog-mongo/core/controller/article_controller"
	"blog-mongo/core/controller/auth_controller"
	"blog-mongo/core/controller/comment_controller"
	"blog-mongo/core/repository/article_repository"
	"blog-mongo/core/repository/auth_repository"
	"blog-mongo/core/repository/comment_respository"
	"blog-mongo/core/service/article_service"
	"blog-mongo/core/service/auth_service"
	"blog-mongo/core/service/comment_service"
	"blog-mongo/helper"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database
var err error
var Mux = mux.NewRouter()
var valid = validator.New()

func init() {
	db,err = helper.Connection()
	helper.PanicCustomException(exception.InternalError{Err:errors.New("terjadi kesalahan pada sistem kami")}, err != nil)
}


func authController() auth_controller.Auth {
	repo := auth_repository.NewAuth()
	service := auth_service.NewAuth(db,repo,valid)
	controller := auth_controller.NewAuth(service)
	return controller
}

func articleController() controller.BasicCrud {
	repo := article_repository.NewArticle()
	service := article_service.NewArticle(valid,db,repo)
	controller := article_controller.NewArticle(service)
	return controller
}

func commentController() controller.BasicCrud {
	repo := comment_respository.NewComment()
	service := comment_service.NewComment(db,valid,repo)
	controller := comment_controller.NewComment(service)
	return controller
}
