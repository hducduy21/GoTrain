package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	router := chi.NewMux()

	router.Get("/", handleExp)
	router.Get("/hducduy21", handleExp)

	listenAddr := os.Getenv("LISTEN_ADD")
	slog.Info("Server is running on port ", listenAddr)
	http.ListenAndServe(listenAddr, router)
}

func handleExp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
