package task

import (
	"context"
	"errors"

	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteSubstep(tid *primitive.ObjectID, sid *primitive.ObjectID, userID *primitive.ObjectID) error {
	var task data.Item
	err := db.Coll.Tasks.FindOne(context.Background(), bson.M{
		"_id":    tid,
		"userId": userID,
	}).Decode(&task)

	if err == mongo.ErrNoDocuments {
		return errors.New(data.TaskNotFound)
	}
	if err != nil {
		return errors.New(data.AnyErrorReadingDB)
	}

	validSubtaskID := false
	for _, substep := range task.SubSteps {
		if *sid == *substep.ID {
			validSubtaskID = true
		}
	}
	if !validSubtaskID {
		return errors.New(data.SubtaskNotFound)
	}

	_, err = db.Coll.Tasks.UpdateOne(context.Background(), bson.M{
		"_id": tid,
	}, bson.M{
		"$pull": bson.M{
			"substeps": bson.M{
				"_id": sid,
			},
		},
	})
	if err != nil {
		return errors.New(data.AnyErrorDeletingSubtask)
	}

	return nil
}
