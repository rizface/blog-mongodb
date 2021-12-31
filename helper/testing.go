package helper

import (
	"blog-mongo/app/model/request"
	"blog-mongo/core/repository/article_repository"
	"blog-mongo/core/repository/comment_respository"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var db,_ = Connection()

func PostArticleTesting() string {
	repo := article_repository.NewArticle()
	result,_ := repo.Post(db,context.Background(),request.Article{
		Author:    "fariz",
		Title:     "JWT",
		Body:      "Json Web Token",
		CreatedAt: time.Now(),
	})
	primitiveObject := result.(*mongo.InsertOneResult).InsertedID
	return primitiveObject.(primitive.ObjectID).Hex()
}

func PostTestingComment() (string,string) {
	claims := TokenValidation(TestingToken()).(*Claims)
	articleId := PostArticleTesting()
	repo := comment_respository.NewComment()
	id,_,_ := repo.Post(db,context.Background(),request.Comment{
		ArticleId: articleId,
		Author:    claims.Username.(string),
		Comment:   "KOMENTAR TESTING",
		CreatedAt: time.Now(),
	})
	return articleId,id
}

func TestingToken() string {
	token,_ := GenerateToken(bson.M{
		"_id": primitive.NewObjectID(),
		"username": "fariz",
		"email": "malfarizzi33@gmail.com",
	})
	return token
}