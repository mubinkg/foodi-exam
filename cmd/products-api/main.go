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
	"github.com/mubinkg/foodi-exam/internal/http/handlers/product"
)

func main() {

	godotenv.Load()

	cfg := config.MustLoadEnv()
	router := http.NewServeMux()

	router.HandleFunc("POST /api/products", product.New())

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
