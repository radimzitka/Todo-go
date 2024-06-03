package user

import (
	"context"
	"errors"

	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Add(user *data.User) (*data.User, error) {
	inserted, err := db.Coll.Users.InsertOne(context.Background(), user)
	if err != nil {
		return nil, errors.New(data.ANY_ERROR_INSERTING_USER)
	}
	iid := inserted.InsertedID.(primitive.ObjectID)
	user.ID = &iid

	return user, nil
}
