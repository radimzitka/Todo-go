package task

import (
	"context"

	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(item *data.Item) (*data.Item, error) {
	// Add thesis to the database
	inserted, err := db.Client.Collection("tasks").InsertOne(context.Background(), item)
	if err != nil {
		return nil, err
	}
	iid := inserted.InsertedID.(primitive.ObjectID)
	item.ID = &iid

	return item, nil
}
