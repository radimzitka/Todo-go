package task

import (
	"context"
	"errors"
	"time"

	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FinishByID(id *primitive.ObjectID, userID *primitive.ObjectID) error {
	var task data.Item
	err := db.Coll.Tasks.FindOne(context.Background(), bson.M{
		"_id":    id,
		"userId": userID,
	}).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return errors.New(data.TaskNotFound)
	}
	if err != nil {
		return err
	}

	if task.Finished {
		return errors.New(data.TaskFinished)
	}

	_, err = db.Coll.Tasks.UpdateOne(context.Background(), bson.M{
		"_id": id,
	}, bson.M{
		"$set": bson.M{
			"finished":     true,
			"timeFinished": time.Now(),
		},
	})
	if err != nil {
		return errors.New(data.AnyErrorDeletingTask)
	}

	return nil
}
