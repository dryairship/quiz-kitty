package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dryairship/quiz-kitty/models"
)

func createUserScore(user string, score int32) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := mongoScoresCollection.InsertOne(ctx, bson.M{
		"_id":   user,
		"score": score,
	})
	if err != nil {
		log.Println("[ERROR] Mongo can't create user score. User: ", user, ", Error: ", err)
	}
}

func GetUserScore(user string) int32 {
	var result models.MongoUserScore
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := mongoScoresCollection.FindOne(ctx, bson.M{"_id": user}).Decode(&result)
	if err != nil {
		log.Println("[ERROR] Mongo can't get user score. User: ", user, ", Error: ", err)
		createUserScore(user, 0)
		return 0
	}
	return result.Score
}

func UpdateUserScore(user string, delta int32) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result := mongoScoresCollection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": user},
		bson.D{primitive.E{
			Key: "$inc",
			Value: bson.D{primitive.E{
				Key:   "score",
				Value: delta,
			}},
		}},
	)
	if result.Err() != nil {
		log.Println("[ERROR] Mongo can't update user score. User: ", user, ", Error: ", result.Err())
		if result.Err() == mongo.ErrNoDocuments {
			createUserScore(user, delta)
		}
	}
}
