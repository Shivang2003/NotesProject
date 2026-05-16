package noteHandler

import (
	"net/http"
	"github.com/Shivang2003/NotesProject/internal/storage"
)

func CreateNoteHandler(storage storage.NotesStorage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// storage.CreateNote()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}

}
