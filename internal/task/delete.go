package task

import (
	"context"
	"errors"

	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Delete(id *primitive.ObjectID, userID *primitive.ObjectID) error {
	result, err := db.Coll.Tasks.DeleteOne(context.Background(), bson.M{
		"_id":    id,
		"userId": userID,
	})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New(data.TASK_NOT_FOUND)
	}
	return nil
}
