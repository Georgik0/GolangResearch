package logger_with_limiter

import (
	"log/slog"
	"os"
	"strconv"
	"testing"
	"time"

	slogmulti "github.com/samber/slog-multi"
	slogsampled "github.com/samber/slog-sampling"
	"github.com/stretchr/testify/assert"
)

type mockIOWriter struct {
	counter int
}

func (w *mockIOWriter) Write(p []byte) (n int, err error) {
	w.counter++

	return len(p), nil
}

func TestSlogSampled(t *testing.T) {
	option := slogsampled.UniformSamplingOption{
		// The sample rate for sampling traces in the range [0.0, 1.0].
		Rate: 0.4, // 40%
	}

	w := &mockIOWriter{}

	for i := 0; i < 1000; i++ {
		sampledLogger := slog.New(
			slogmulti.Pipe(option.NewMiddleware()).
				Handler(slog.NewJSONHandler(w, nil)),
		)
		sampledLogger.Warn(strconv.Itoa(i))
	}

	result40percentOfLimit := w.counter > 300 && w.counter < 500

	assert.True(t, result40percentOfLimit)
	slog.Info(strconv.Itoa(w.counter))
}

func TestFormat(t *testing.T) {
	const durationLoggingRate = 0.01

	option := slogsampled.UniformSamplingOption{
		Rate: durationLoggingRate,
	}

	logger := slog.New(
		slogmulti.Pipe(option.NewMiddleware()).
			Handler(slog.NewJSONHandler(os.Stdout, nil)),
	)

	for i := 0; i < 1000; i++ {
		logger.Error("test msg ... test", "duration_ms", 5*time.Millisecond)
	}
}
