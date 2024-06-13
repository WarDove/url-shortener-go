package main

import (
	"fmt"
	"url-shortener-go/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg) // TODO: remove debug print
	// TODO: init logger: slog
	// TODO: init db: sqlite
	// TODO: init router: chi, chi render
	// TODO: run server
}
