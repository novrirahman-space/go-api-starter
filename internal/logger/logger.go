package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func New(env string) zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log := zerolog.New(output).With().Timestamp().Logger()
	if env == "production" {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	}
	return log
}
