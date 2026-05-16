package userHandler

import (
	"encoding/json"
	"net/http"

	"github.com/Shivang2003/NotesProject/internal/storage"
	"github.com/Shivang2003/NotesProject/internal/types"
)

func CreateUserHandler(storage storage.UserStorage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// storage.CreateUser()
		
		var user types.User

		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		storage.CreateUser(user)

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(
			map[string]string{"message": "User created successfully"},
		)
	}

}
