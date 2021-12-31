package auth_repository

import (
	"blog-mongo/app/model/request"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type auth struct{}

func NewAuth() Auth{
	return &auth{}
}

func (a *auth) Register(db *mongo.Database, ctx context.Context,request request.Register) bool {
	result,err := db.Collection("users").InsertOne(ctx,bson.D{{"username",request.Username},{"email",request.Email},{"password",request.Password},{"createdAt",time.Now()}})
	if err != nil || result == nil{
		return false
	}
	return true
}

func (a *auth) Login(db *mongo.Database, ctx context.Context, email string) (bson.M,error) {
	var result bson.M
	cursor := db.Collection("users").FindOne(ctx,bson.M{"email":bson.M{"$eq":email}})
	err := cursor.Decode(&result)
	return result,err
}





