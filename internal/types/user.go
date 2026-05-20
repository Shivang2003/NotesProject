package types

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	ID       bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string        `json:"username" bson:"username" validate:"required"`
	Email    string        `json:"email" bson:"email" validate:"required,email"`
	Password string        `json:"password" bson:"password" validate:"required"`

	Notes []Notes `json:"notes" bson:"notes"`
}
