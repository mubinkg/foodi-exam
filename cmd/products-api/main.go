package main

import (
	"log/slog"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/mubinkg/foodi-exam/internal/config"
)

func main() {

	godotenv.Load()

	config.MustLoadEnv()
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Wellcome to foodi crud"))
	})

	server := http.Server{
		Addr:    "localhost:8082",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		slog.Warn("Server not started")
	}
}
