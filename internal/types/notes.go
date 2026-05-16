package types

type Notes struct {
	content string `json:"content" bson:"content"`
	createdAt string `json:"created_at" bson:"created_at"`
	updatedAt string `json:"updated_at" bson:"updated_at"`
}