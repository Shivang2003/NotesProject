package types

import "go.mongodb.org/mongo-driver/v2/bson"


type Notes struct {
	ID        bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Content   string             `json:"content" bson:"content"`
	CreatedAt string             `json:"created_at" bson:"created_at"`
	UpdatedAt string             `json:"updated_at" bson:"updated_at"`
}
