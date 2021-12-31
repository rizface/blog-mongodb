package comment_respository

import (
	"blog-mongo/app/model/request"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type comment struct{}

func NewComment() Comment {
	return &comment{}
}

func (c comment) Post(db *mongo.Database, ctx context.Context, request request.Comment) (string,bool,error){
	objId, _ := primitive.ObjectIDFromHex(request.ArticleId)
	commentId := primitive.NewObjectID()
	result,err := db.Collection("articles").UpdateOne(ctx, bson.M{
		"_id": objId,
	}, bson.M{"$push": bson.M{"comments": bson.M{
		"id": commentId,
		"author": request.Author, "comment": request.Comment,"createdAt":request.CreatedAt,
	}}})
	return commentId.Hex(),result.ModifiedCount > 0, err
}

func (c comment) Delete(db *mongo.Database, ctx context.Context, request request.DeleteComment) (bool, error) {
	articleId,_ := primitive.ObjectIDFromHex(request.ArticleId)
	commentId,_ := primitive.ObjectIDFromHex(request.CommentId)
	result,err := db.Collection("articles").UpdateOne(ctx,bson.M{
		"$and":bson.A{bson.M{"_id":bson.M{"$eq":articleId}},bson.M{
			"$or":bson.A{bson.M{"author":request.Username},bson.M{"comments.author":request.Username}},
		}},
	},bson.M{
		"$pull":bson.M{
			"comments":bson.M{
				"id":commentId,
			},
		},
	})
	return result.ModifiedCount > 0, err
}

func (c comment) Update(db *mongo.Database, ctx context.Context, request request.UpdateComment) (bool, error) {
	articleId,_ := primitive.ObjectIDFromHex(request.ArticleId)
	commentId,_ := primitive.ObjectIDFromHex(request.CommentId)
	result,err := db.Collection("articles").UpdateOne(ctx,bson.M{
		"$and":bson.A{bson.M{"_id":articleId},bson.M{"comments.id":commentId},bson.M{"comments.author": request.Username,}},
	},bson.M{
		"$set":bson.M{
			"comments.$.comment":request.Comment,
		},
	})
	return result.ModifiedCount > 0, err
}

