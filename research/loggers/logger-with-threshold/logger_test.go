package logger_with_threshold

import (
	"context"
	"io"
	"log/slog"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Skip(
		"Тест проверяет механизм основанный на тикере, " +
			"поэтому он выключен, чтобы не тормозить время прохождения тестов",
	)

	var messageCount atomic.Int64

	testHandler := slog.NewTextHandler(io.Discard, &slog.HandlerOptions{})
	wrappedHandler := &countingHandler{
		handler: testHandler,
		count:   &messageCount,
	}

	handler := New(wrappedHandler, 5, 2*time.Second)
	logger := slog.New(handler)
	defer handler.Stop()

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			logger.Info("test message")
			wg.Done()
		}()
	}

	wg.Wait()

	count := messageCount.Load()
	assert.Equal(t, int64(5), count)

	// Ждем сброса счетчика и проверяем, что логирование возобновляется
	time.Sleep(3 * time.Second)
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			logger.Info("test message after reset")
			wg.Done()
		}()
	}

	wg.Wait()

	count = messageCount.Load()
	assert.Equal(t, int64(8), count)
}

type countingHandler struct {
	handler slog.Handler
	count   *atomic.Int64
}

func (h *countingHandler) Handle(ctx context.Context, r slog.Record) error {
	h.count.Add(1)
	return h.handler.Handle(ctx, r)
}

func (h *countingHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &countingHandler{handler: h.handler.WithAttrs(attrs), count: h.count}
}

func (h *countingHandler) WithGroup(name string) slog.Handler {
	return &countingHandler{handler: h.handler.WithGroup(name), count: h.count}
}

func (h *countingHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}
