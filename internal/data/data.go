package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ANY_ERROR_READING_DTB       = "error reading from database"
	ANY_ERROR_INSERTING_TASK    = "error while inserting new task"
	ANY_ERROR_INSERTING_SUBTASK = "error while inserting new task"
	ANY_ERROR_DELETING_TASK     = "error while searching for the task"
	ANY_ERROR_DELETING_SUBTASK  = "error while searching for the subtask"
	TASK_NOT_FOUND              = "task was not found"
	SUBTASK_NOT_FOUND           = "subtask was not found"
	SUBTASK_FINISHED            = "subtask already finished"
	TASK_FINISHED               = "task already finished"
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
