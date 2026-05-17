package userHandler

import (
	"encoding/json"
	"errors"
	"net/http"
	"io"
	"github.com/Shivang2003/NotesProject/internal/storage"
	"github.com/Shivang2003/NotesProject/internal/types"
	"github.com/go-playground/validator/v10"
	"github.com/Shivang2003/NotesProject/internal/utils/response"

)

func CreateUserHandler(storage storage.UserStorage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// storage.CreateUser()
		
		var user types.User

		err := json.NewDecoder(r.Body).Decode(&user)

		if errors.Is(err, io.EOF){
			http.Error(w, "Request body cannot be empty", http.StatusBadRequest)
			return
		} 

		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		validate := validator.New()

		if err := validate.Struct(user); err != nil {

			validateErrs := err.(validator.ValidationErrors)

			response.WriteJson(w, http.StatusBadRequest,response.VlidateError(validateErrs))
			return
		}

		storage.CreateUser(user)

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(
			map[string]string{"message": "User created successfully"},
		)
	}

}
