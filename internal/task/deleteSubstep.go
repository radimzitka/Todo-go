package task

import (
	"context"
	"errors"

	"github.com/radimzitka/zitodo-mongo/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteSubstep(tid *primitive.ObjectID, sid *primitive.ObjectID) error {
	// Overovat, jestli ID v DTB existuje nebo jak to dělat? Já takhle nepoznám chybu...
	delete, err := db.Coll.Tasks.UpdateOne(context.Background(), bson.M{
		"_id": tid,
	}, bson.M{
		"$pull": bson.M{
			"substeps": bson.M{
				"_id": sid,
			},
		},
	})
	if err != nil {
		return errors.New("error occured while deleting subtask")
	}

	if delete.MatchedCount == 0 {
		return errors.New("subtask not found")
	}

	return nil
}
