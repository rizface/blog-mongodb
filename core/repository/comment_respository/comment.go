package comment_respository

import (
	"blog-mongo/app/model/request"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Comment interface {
	Post(db *mongo.Database, ctx context.Context, request request.Comment) (string,bool,error)
	Delete(db *mongo.Database, ctx context.Context, request request.DeleteComment) (bool,error)
	Update(db *mongo.Database, ctx context.Context, request request.UpdateComment) (bool,error)
}
