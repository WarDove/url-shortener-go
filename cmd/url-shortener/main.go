package main

import (
	"log/slog"
	"os"
	"url-shortener-go/internal/config"
	"url-shortener-go/internal/lib/logger/sl"
	"url-shortener-go/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "development"
	envProd  = "production"
)

func setupLogger(env string) *slog.Logger {

	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("Starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("Debug messages are enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("Failed to initialize storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	//id, err := storage.SaveURL("https://google.com", "google")
	//if err != nil {
	//	log.Error("Failed to save url", sl.Err(err))
	//	os.Exit(1)
	//}
	//
	//log.Info("saved url", slog.Int64("id", id))

	//str, err := storage.GetURL("google")
	//if err != nil {
	//	log.Error("Failed to get url", sl.Err(err))
	//	os.Exit(1)
	//}
	//fmt.Println(str, "test")
	//
	//if storage.DeleteURL("google") != nil {
	//	log.Error("Failed to delete url", sl.Err(err))
	//	os.Exit(1)
	//}

	//id, err = storage.SaveURL("https://google.az", "google")
	//if err != nil {
	//	log.Error("Failed to save url", sl.Err(err))
	//	os.Exit(1)
	//}
	//
	//log.Info("saved url", slog.Int64("id", id))

	// TODO: init router: chi, chi render
	// TODO: run server
}
