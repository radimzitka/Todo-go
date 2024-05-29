package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SubStep struct {
	ID           *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title        string              `json:"title" bson:"title"`
	Done         bool                `json:"done" bson:"done"`
	FinishedTime *time.Time          `json:"finishedTime" bson:"finishedTime"`
}

type Item struct {
	ID           *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title        string              `json:"title" bson:"title"`
	Description  string              `json:"description" bson:"description"`
	NotToJson    string              `json:"-" bson:"-"`
	TimeAdded    *time.Time          `json:"timeAdded" bson:"timeAdded"`
	SubSteps     []*SubStep          `json:"substeps" bson:"substeps"`
	Finished     bool                `json:"finished" bson:"finished"`
	TimeFinished *time.Time          `json:"timeFinished" bson:"timeFinished"`
}
