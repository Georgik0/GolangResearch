package main

import (
	"log/slog"
	"os"
	"time"

	slogmulti "github.com/samber/slog-multi"
	slogsampled "github.com/samber/slog-sampling"
)

func main() {
	// Создаем logger с ограничением: максимум 4 лога в секунду
	option := slogsampled.AbsoluteSamplingOption{
		Tick: 1 * time.Second, // интервал времени
		Max:  4,               // максимум 4 лога за интервал
	}

	logger := slog.New(
		slogmulti.Pipe(option.NewMiddleware()).
			Handler(slog.NewJSONHandler(os.Stdout, nil)),
	)

	slog.Info("=== Попытка залогировать 10 сообщений быстро ===")
	slog.Info("Ожидаемый результат: только первые 4 сообщения пройдут в течение секунды\n")

	// Попытка залогировать 10 сообщений быстро
	for i := 0; i < 10; i++ {
		logger.Info("Сообщение", "номер", i, "время", time.Now().Format("15:04:05.000"))
	}

	// Ждем 1 секунду, чтобы лимит сбросился
	time.Sleep(1 * time.Second)

	slog.Info("\n=== Через секунду лимит сбросился, логируем еще 3 сообщения ===\n")

	// Теперь можем снова залогировать до 4 сообщений
	for i := 10; i < 13; i++ {
		logger.Info("Сообщение после сброса", "номер", i, "время", time.Now().Format("15:04:05.000"))
	}
}
