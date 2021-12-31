package setup

import (
	"blog-mongo/app/middleware"
	"blog-mongo/route"
	"github.com/gorilla/mux"
	"net/http"
)

func AuthRouter() *mux.Router {
	controller := authController()
	auth := Mux.NewRoute().Subrouter()
	auth.Use(middleware.ErrorHandler)
	auth.HandleFunc(route.REGISTER,controller.Register).Methods(http.MethodPost)
	auth.HandleFunc(route.LOGIN,controller.Login).Methods(http.MethodPost)
	return auth
}

func ArticleRouter() *mux.Router {
	controller := articleController()
	article := Mux.NewRoute().Subrouter()
	article.Use(middleware.ErrorHandler,middleware.Auth)

	article.HandleFunc(route.ARTIKEL,controller.Get).Methods(http.MethodGet)
	article.HandleFunc(route.ARTIKEL,controller.Post).Methods(http.MethodPost)
	article.HandleFunc(route.ARTIKEL_MANIPULATION,controller.Delete).Methods(http.MethodDelete)
	article.HandleFunc(route.ARTIKEL_MANIPULATION,controller.Update).Methods(http.MethodPut)

	return article
}

func CommentRouter() *mux.Router {
	controller := commentController()
	comment := Mux.NewRoute().Subrouter()
	comment.Use(middleware.ErrorHandler,middleware.Auth)

	comment.HandleFunc(route.COMMENT,controller.Post).Methods(http.MethodPost)
	comment.HandleFunc(route.MANIPULATION_COMMENT,controller.Delete).Methods(http.MethodDelete)
	comment.HandleFunc(route.MANIPULATION_COMMENT,controller.Update).Methods(http.MethodPut)
	return comment
}