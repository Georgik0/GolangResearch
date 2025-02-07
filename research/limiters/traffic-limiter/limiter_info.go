package trafficlimiter

import "time"

type limiterKey string

const (
	ExampleKey     = "example_..."
	ExampleTimeout = time.Minute
	ExampleLimit   = 1
)
