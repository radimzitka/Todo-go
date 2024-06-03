package task

import (
	"context"
	"errors"

	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateSubstep(substep *data.SubStep, id *primitive.ObjectID) (*data.Item, error) {
	var task data.Item
	opts := options.FindOneAndUpdate()
	opts.SetReturnDocument(options.After)

	sid := primitive.NewObjectID()
	substep.ID = &sid
	err := db.Coll.Tasks.FindOneAndUpdate(context.Background(), bson.M{
		"_id": id,
	}, bson.M{
		"$push": bson.M{
			"substeps": substep,
		},
	}, opts).Decode(&task)

	if err == mongo.ErrNoDocuments {
		return nil, errors.New(data.TASK_NOT_FOUND)
	}
	if err != nil {
		return nil, errors.New(data.ANY_ERROR_INSERTING_SUBTASK)
	}

	return &task, nil
}
