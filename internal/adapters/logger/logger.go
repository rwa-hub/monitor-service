package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Log = zerolog.New(os.Stdout).With().
	Timestamp().
	Caller().
	Str("service", "monitor-service").
	Logger()

func Init() {
	zerolog.TimeFieldFormat = time.RFC3339
}
