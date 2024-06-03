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

func DeleteSubstep(tid *primitive.ObjectID, sid *primitive.ObjectID) error {
	var task data.Item
	err := db.Coll.Tasks.FindOne(context.Background(), bson.M{
		"_id": tid,
	}).Decode(&task)

	if err == mongo.ErrNoDocuments {
		return errors.New(data.TASK_NOT_FOUND)
	}
	if err != nil {
		return errors.New(data.ANY_ERROR_READING_DTB)
	}

	validSubtaskID := false
	for _, substep := range task.SubSteps {
		if *sid == *substep.ID {
			validSubtaskID = true
		}
	}
	if !validSubtaskID {
		return errors.New(data.SUBTASK_NOT_FOUND)
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
		return errors.New(data.ANY_ERROR_DELETING_SUBTASK)
	}

	return nil
}
