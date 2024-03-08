package main

import (
	"log/slog"
	"logging"
)

func main() {
	logging.Initialize("CLI")
	slog.Info("CLI Application")
}
