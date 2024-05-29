package task

import (
	"context"

	"github.com/radimzitka/zitodo-mongo/internal/data"
	"github.com/radimzitka/zitodo-mongo/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateSubstep(substep *data.SubStep, id *primitive.ObjectID) (*data.Item, error) {
	// Add thesis to the database
	// Jaky je postup? 1. Nactu si item, do ktereho chci vkladat, 2. vlozim novy zaznam v subtascich a 3. updatuju ho
	var task data.Item
	opts := options.FindOneAndUpdate()
	opts.SetReturnDocument(options.After)

	sid := primitive.NewObjectID()
	substep.ID = &sid
	err := db.Client.Collection("tasks").FindOneAndUpdate(context.Background(), bson.M{
		"_id": id,
	}, bson.M{
		"$push": bson.M{
			"substeps": substep,
		},
	}, opts).Decode(&task)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

// zapamatuj si: $push
