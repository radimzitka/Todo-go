package user

import (
	"context"
	"errors"

	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteByID(id *primitive.ObjectID) error {
	result, err := db.Coll.Users.DeleteOne(context.Background(), bson.M{
		"_id": id,
	})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New(data.USER_NOT_FOUND)
	}
	return nil
}