package trafficlimiter

import (
	"time"
)

const (
	defaultLimit   = 1
	defaultTimeout = time.Hour
)

type Limiter struct {
	timeout time.Duration
	counter int32
	limit   int32
	ch      chan struct{}
	out     chan struct{}
}

func StartNewLimiter(
	newTimeuot time.Duration,
	newLimit int32,
) *Limiter {
	timeout := newTimeuot
	if timeout == 0 {
		timeout = defaultTimeout
	}

	limit := newLimit
	if limit == 0 {
		limit = defaultLimit
	}

	l := &Limiter{
		timeout: timeout,
		limit:   limit,
		ch:      make(chan struct{}),
		out:     make(chan struct{}),
	}

	l.run()

	return l
}

func (l *Limiter) Cancel() {
	if l == nil || l.out == nil {
		return
	}

	l.out <- struct{}{}

	close(l.out)
}

func (l *Limiter) run() {
	if l == nil || l.timeout == 0 {
		return
	}

	go func() {
		for {
			if l.counter >= l.limit {
				time.Sleep(l.timeout)

				l.counter = 0
			}

			select {
			case _ = <-l.ch:
			case _ = <-l.out:
				return
			}

			l.counter++

		}
	}()
}

func (l *Limiter) Allow() bool {
	if l == nil {
		return false
	}

	select {
	case l.ch <- struct{}{}:
		return true
	default:
		return false
	}
}
