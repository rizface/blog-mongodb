package comment_respository

import (
	"blog-mongo/app/model/request"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Comment interface {
	GetById(db *mongo.Database, ctx context.Context, articleId string, commentId string) (bson.M,error)
	Post(db *mongo.Database, ctx context.Context, request request.Comment) (string,bool,error)
	Delete(db *mongo.Database, ctx context.Context, request request.DeleteComment) (bool,error)
	Update(db *mongo.Database, ctx context.Context, request request.UpdateComment) (bool,error)
}
