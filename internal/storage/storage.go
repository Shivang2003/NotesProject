package storage

import (
	"github.com/Shivang2003/NotesProject/internal/types"
)

type UserStorage interface {
	CreateUser(user types.User) error
	// DeleteUser(id string) error
	UpdateUser(id string) error
	// ReadUser(id string)
}

type NotesStorage interface {
	// CreateNote(note types.Notes)
	// DeleteNote(id string)
	// UpdateNote(id string)
	// ReadNote(id string)
}
