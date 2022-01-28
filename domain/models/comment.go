package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	Id      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Message string             `json:"message" bson:"message,omitempty"`
}
