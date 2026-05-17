package types

type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username" validate:"required"`
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required"`

	Notes []Notes `json:"notes" bson:"notes"`
}