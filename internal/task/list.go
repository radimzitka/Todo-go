package task

import (
	"context"
	"errors"

	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/db"
	"go.mongodb.org/mongo-driver/bson"
)

func List() ([]*data.Item, error) {
	cursor, err := db.Coll.Tasks.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.New("error during access to dtb")
	}

	tasks := make([]*data.Item, 0)
	for cursor.Next(context.Background()) {
		var t data.Item
		cursor.Decode(&t)
		tasks = append(tasks, &t)
	}

	return tasks, nil
}
