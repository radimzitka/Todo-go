package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// opravit
	AnyErrorReadingDB        = "error reading from database"
	AnyErrorInsertingTask    = "error while inserting new task"
	AnyErrorInsertingSubtask = "error while inserting new task"
	AnyErrorDeletingTask     = "error while searching for the task"
	AnyErrorDeletingSubtask  = "error while searching for the subtask"
	TaskNotFound             = "task was not found"
	SubtaskNotFound          = "subtask was not found"
	SubtaskFinished          = "subtask already finished"
	TaskFinished             = "task already finished"
	AnyErrorInsertingUser    = "error while inserting new user"
	UserNotFound             = "user not found"
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
	UserID       *primitive.ObjectID `json:"userId" bson:"userId"`
}

type User struct {
	ID       *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string              `json:"name" bson:"name"`
	Email    string              `json:"email" bson:"email"`
	Password string              `json:"-" bson:"password"`
}
