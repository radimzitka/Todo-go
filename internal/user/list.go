package user

import (
	"context"
	"errors"

	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/db"
	"go.mongodb.org/mongo-driver/bson"
)

func List() ([]*data.User, error) {
	cursor, err := db.Coll.Users.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.New(data.AnyErrorReadingDB)
	}

	users := make([]*data.User, 0)
	for cursor.Next(context.Background()) {
		var u data.User
		cursor.Decode(&u)
		users = append(users, &u)
	}

	return users, nil
}
