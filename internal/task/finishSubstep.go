package task

import (
	"context"
	"errors"
	"time"

	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FinishSubstep(tid *primitive.ObjectID, sid *primitive.ObjectID) (*data.SubStep, error) {
	var task data.Item
	opts := options.FindOneAndUpdate()
	opts.SetReturnDocument(options.After)

	err := db.Client.Collection("tasks").FindOne(context.Background(), bson.M{
		"_id":          tid,
		"substeps._id": sid,
	}).Decode(&task)

	if err != nil {
		return nil, err
	}

	// Is substep already done?
	for _, substep := range task.SubSteps {
		if *substep.ID == *sid {
			if substep.Done {
				return nil, errors.New("substep is already done")
			}
		}
	}

	err = db.Client.Collection("tasks").FindOneAndUpdate(context.Background(), bson.M{
		"_id":          tid,
		"substeps._id": sid,
	}, bson.M{
		"$set": bson.M{
			"substeps.$.done":         true,
			"substeps.$.finishedTime": time.Now(),
		},
	}, opts).Decode(&task)

	if err != nil {
		return nil, err
	}

	for _, substep := range task.SubSteps {
		if *substep.ID == *sid {
			return substep, nil
		}
	}

	return nil, nil
}
