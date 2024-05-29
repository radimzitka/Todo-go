package task

import (
	"context"
	"errors"
	"time"

	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FinishByID(id *primitive.ObjectID) error {
	var task data.Item
	err := db.Client.Collection("tasks").FindOne(context.Background(), bson.M{
		"_id": id,
	}).Decode(&task)
	if err != nil {
		return err
	}

	if task.Finished {
		return errors.New("task was already finished")
	}

	_, err = db.Client.Collection("tasks").UpdateOne(context.Background(), bson.M{
		"_id": id,
	}, bson.M{
		"$set": bson.M{
			"finished":     true,
			"timeFinished": time.Now(),
		},
	})

	return err
}