package task

import (
	"context"
	"errors"

	"github.com/radimzitka/zitodo-mongo/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteByID(id *primitive.ObjectID) error {
	result, err := db.Client.Collection("tasks").DeleteOne(context.Background(), bson.M{
		"_id": id,
	})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("try to delete non-exist task")
	}
	return nil
}