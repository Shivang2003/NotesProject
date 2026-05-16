package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Shivang2003/NotesProject/internal/http/handlers/noteHandler"
	"github.com/Shivang2003/NotesProject/internal/http/handlers/userHandler"
	"github.com/Shivang2003/NotesProject/storage/mongodb"
	"github.com/joho/godotenv"
)

func main() {

	//load config

	//load env variables

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}

	//db setup

	store, err := mongodb.Connect()

	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v\n", err)
		return
	}

	//setup router

	router := http.NewServeMux()

	//user route apis

	router.HandleFunc("POST /api/v1/users/createUser", userHandler.CreateUserHandler(store))

	//notes route apis
	router.HandleFunc("POST /api/v1/notes/createNote", noteHandler.CreateNoteHandler(store))

	//setup server

	server := http.Server{
		Addr:    "localhost:8082",
		Handler: router,
	}

	fmt.Printf("Server started at %s", server.Addr)
	//graceful shutdown

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// slog.Info("Started server at ", slog.String("address", cfg.HTTPServer.Addr))
	// fmt.Printf("Server started at %s", cfg.HTTPServer.Addr)

	go func() { //separate goroutine to start the server and listen for incoming requests which allows the main goroutine to listen for shutdown signals and perform graceful shutdown when needed
		err := server.ListenAndServe() //this
		if err != nil {
			log.Fatal("failed to start server", err.Error())
		}
	}()

	<-done

	slog.Info("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = server.Shutdown(ctx)

	if err != nil {
		slog.Error("failed to shutdown server", "error", err.Error())
	}

	slog.Info("Server stopped gracefully")

}
