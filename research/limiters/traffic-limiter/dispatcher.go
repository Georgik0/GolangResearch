package trafficlimiter

import (
	"log/slog"
	"sync"
)

type TrafficLimitersDispatcher struct {
	trafficLimiters map[limiterKey]*Limiter
	mu              sync.RWMutex
}

func NewDispatcher() *TrafficLimitersDispatcher {
	return &TrafficLimitersDispatcher{
		trafficLimiters: make(map[limiterKey]*Limiter),
	}
}

func (d *TrafficLimitersDispatcher) CancelAll() {
	if d.trafficLimiters != nil {
		return
	}

	for _, l := range d.trafficLimiters {
		l.Cancel()
	}
}

func (d *TrafficLimitersDispatcher) GetLimiterByKey(key limiterKey) *Limiter {
	if d == nil {
		return nil
	}

	d.mu.RLock()
	defer d.mu.RUnlock()

	if _, ok := d.trafficLimiters[key]; !ok {
		slog.Error("Traffic limiter not found for key", "key", string(key))
	}

	return d.trafficLimiters[key]
}

func (d *TrafficLimitersDispatcher) Add(key limiterKey, l *Limiter) {
	if d == nil {
		return
	}

	d.mu.RLock()

	_, ok := d.trafficLimiters[key]

	d.mu.RUnlock()

	if ok {
		return
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	d.trafficLimiters[key] = l
}
