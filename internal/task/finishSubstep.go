package task

import (
	"context"
	"errors"
	"time"

	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const SubtaskFinished = "task already finished"

func FinishSubstep(tid *primitive.ObjectID, sid *primitive.ObjectID, userID *primitive.ObjectID) (*data.SubStep, error) {
	var task data.Item
	opts := options.FindOneAndUpdate()
	opts.SetReturnDocument(options.After)

	err := db.Coll.Tasks.FindOne(context.Background(), bson.M{
		"_id":    tid,
		"userId": userID,
	}).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New(data.TaskNotFound)
	}
	if err != nil {
		return nil, errors.New(data.AnyErrorDeletingTask)
	}

	// Is substep already done?
	for _, substep := range task.SubSteps {
		if *substep.ID == *sid {
			if substep.Done {
				return nil, errors.New(data.SubtaskFinished)
			}
		}
	}

	err = db.Coll.Tasks.FindOneAndUpdate(context.Background(), bson.M{
		"_id":          tid,
		"substeps._id": sid,
	}, bson.M{
		"$set": bson.M{
			"substeps.$.done":         true,
			"substeps.$.finishedTime": time.Now(),
		},
	}, opts).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New(data.SubtaskNotFound)
	}
	if err != nil {
		return nil, errors.New(data.AnyErrorDeletingSubtask)
	}

	for _, substep := range task.SubSteps {
		if *substep.ID == *sid {
			return substep, nil
		}
	}

	return nil, nil
}
