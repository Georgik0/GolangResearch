package logger_with_threshold

import (
	"context"
	"log/slog"
	"sync/atomic"
	"time"
)

type ThresholdHandler struct {
	handler       slog.Handler
	counter       atomic.Int64
	threshold     int64
	resetInterval time.Duration
	ticker        *time.Ticker
	stopChan      chan struct{}
}

func New(
	handler slog.Handler,
	threshold int64,
	resetInterval time.Duration,
) *ThresholdHandler {
	h := &ThresholdHandler{
		handler:       handler,
		threshold:     threshold,
		stopChan:      make(chan struct{}),
		resetInterval: resetInterval,
	}

	if resetInterval > 0 {
		h.ticker = time.NewTicker(resetInterval)
		go h.periodicReset()
	}

	return h
}

func (h *ThresholdHandler) periodicReset() {
	for {
		select {
		case <-h.ticker.C:
			h.Reset()
		case <-h.stopChan:
			h.ticker.Stop()
			return
		}
	}
}

func (h *ThresholdHandler) Stop() {
	if h.ticker != nil {
		close(h.stopChan)
	}
}

func (h *ThresholdHandler) Reset() {
	h.counter.Store(0)
}

func (h *ThresholdHandler) Handle(ctx context.Context, r slog.Record) error {
	count := h.counter.Add(1)
	if count <= h.threshold {
		return h.handler.Handle(ctx, r)
	}
	return nil
}

func (h *ThresholdHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h.Stop()

	return New(h.handler.WithAttrs(attrs), h.threshold, h.resetInterval)
}

func (h *ThresholdHandler) WithGroup(name string) slog.Handler {
	h.Stop()

	return New(h.handler.WithGroup(name), h.threshold, h.resetInterval)
}

func (h *ThresholdHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}
