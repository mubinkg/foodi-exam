package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/mubinkg/foodi-exam/internal/config"
)

func main() {

	godotenv.Load()

	cfg := config.MustLoadEnv()
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Wellcome to foodi crud"))
	})

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	slog.Info("Server started on", slog.String("address", cfg.Address))

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			slog.Warn("Server not started")
		}
	}()
	<-done

	slog.Info("Shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*1000)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown the server")
	}
	slog.Info("server shut down successfully")
}
