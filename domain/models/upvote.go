package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Upvote struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatedAt uint64             `json:"created_at" bson:"created_at,omitempty"`
	Type      string             `json:"type" bson:"type,omitempty"`
	CommentId primitive.ObjectID `json:"comment_id" bson:"comment_id,omitempty"`
}
