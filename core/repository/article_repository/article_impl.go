package article_repository

import (
	"blog-mongo/app/model/request"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type article struct{}

func NewArticle() Article {
	return &article{}
}

func (a *article) Get(db *mongo.Database, ctx context.Context, skip int64, limit int64) (*mongo.Cursor,error) {
	cursor,err := db.Collection("articles").Find(ctx,bson.M{},options.Find().SetSkip(skip).SetLimit(limit).SetSort(bson.D{{"createdAt",-1}}))
	return cursor,err
}

func (a *article) GetById(db *mongo.Database, ctx context.Context, articleId string) *mongo.SingleResult {
	objId,_ := primitive.ObjectIDFromHex(articleId)
	cursor := db.Collection("articles").FindOne(ctx,bson.M{"_id":objId},options.FindOne().SetProjection(bson.M{
		"title":1,
		"body":1,
		"createdAt":1,
	}))
	return cursor
}

func (a *article) Post(db *mongo.Database, ctx context.Context, request request.Article) (interface{},error)  {
	result,err := db.Collection("articles").InsertOne(ctx,bson.M{
		"author":request.Author,
		"body":request.Body,
		"title":request.Title,
		"createdAt":time.Now(),
	})
	return result,err
}

func (a *article) Delete(db *mongo.Database, ctx context.Context, author string,articleId string) (bool,error){
	objectId,_ := primitive.ObjectIDFromHex(articleId)
	result,err := db.Collection("articles").DeleteOne(ctx,bson.D{{"$and",bson.A{
		bson.D{{"_id", objectId}},
		bson.D{{"author", author}},
	}}})
	return result.DeletedCount > 0, err
}

func (a *article) Update(db *mongo.Database, ctx context.Context, request request.UpdateArticle) (bool, error) {
	objectId,_ := primitive.ObjectIDFromHex(request.Id)
	result,err := db.Collection("articles").UpdateOne(ctx,bson.D{{"$and", bson.A{
		bson.M{"author": request.Author},
		bson.M{"_id":objectId},
	}}},bson.M{
		"$set":bson.M{
			"title": request.Title,
			"body": request.Body,
		},
	})
	return result.ModifiedCount > 0,err
}




