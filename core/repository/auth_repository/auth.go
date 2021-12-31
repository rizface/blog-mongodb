package auth_repository

import (
	"blog-mongo/app/model/request"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Auth interface {
	Register(db *mongo.Database,ctx context.Context,request request.Register) bool
	Login(db *mongo.Database, ctx context.Context, email string) (bson.M,error)
}
