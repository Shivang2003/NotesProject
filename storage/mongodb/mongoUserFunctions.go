package mongodb

import (
	"context"
	"errors"

	"github.com/Shivang2003/NotesProject/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

func (m *MongoStore) checkUserExists(user types.User) (bool, error) {

	var existingUser types.User

	err := m.UsersCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)

	if err == nil {
		return false, errors.New("User with this email already exists")
	}

	if err != mongo.ErrNoDocuments {
		return false, err
	}

	return true, nil
}

func (m *MongoStore) CreateUser(user types.User) error {

	isUnique, err := m.checkUserExists(user)

	if !isUnique {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	_, err = m.UsersCollection.InsertOne(context.TODO(), user)

	if err != nil {
		return err
	}
	return nil
}

func (m *MongoStore) UpdateUser(id string) error {
	isUnique, err := m.checkUserExists(types.User{ID: id})

	if !isUnique {
		return err
	}

	return nil

}
