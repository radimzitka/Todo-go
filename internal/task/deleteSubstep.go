package task

import (
	"context"
	"errors"

	"github.com/radimzitka/zitodo-mongo/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteSubstep(tid *primitive.ObjectID, sid *primitive.ObjectID) error {
	delete, err := db.Client.Collection("tasks").UpdateOne(context.Background(),
		bson.M{"_id": tid, "substeps._id": sid}, bson.M{
			"$pull": bson.M{
				"substeps": bson.M{
					"_id": sid,
				},
			},
		},
	)
	if err != nil {
		return err
	}

	if delete.MatchedCount == 0 {
		return errors.New("trying to delete non-exist substep")
	}

	return nil
}
