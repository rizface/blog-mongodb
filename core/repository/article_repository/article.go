package article_repository

import (
	"blog-mongo/app/model/request"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Article interface {
	Get(db *mongo.Database, ctx context.Context, skip int64 ,limit int64) (*mongo.Cursor,error)
	Post(db *mongo.Database, ctx context.Context,request request.Article) (interface{},error)
	Delete(db *mongo.Database, ctx context.Context, author string,articleId string) (bool,error)
	Update(db *mongo.Database, ctx context.Context, request request.UpdateArticle) (bool,error)
}
