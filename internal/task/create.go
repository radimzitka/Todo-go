package task

import (
	"context"
	"errors"

	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(item *data.Item) (*data.Item, error) {
	inserted, err := db.Coll.Tasks.InsertOne(context.Background(), item)
	if err != nil {
		return nil, errors.New(data.AnyErrorInsertingTask)
	}
	iid := inserted.InsertedID.(primitive.ObjectID)
	item.ID = &iid

	return item, nil
}
