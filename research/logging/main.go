package main

import (
	"log/slog"
)

func main() {
	slog.Warn("[advCache] requestedID too large",
		"lastAdvID", 1,
		"requestedID", 1,
	)
}
