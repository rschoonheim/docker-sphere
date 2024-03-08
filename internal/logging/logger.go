package logging

import (
	"log/slog"
	"os"
)

// Initialize - Configures the slog logger to a standard.
func Initialize(moduleName string) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
	logger = logger.With("module", moduleName)
	slog.SetDefault(logger)
}
