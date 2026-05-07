package main

import (
	"log/slog"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	// Seconds field, required
	// cron.New(cron.WithSeconds())

	// Создаем новый экземпляр cron
	// Seconds field, optional
	c := cron.New(cron.WithParser(cron.NewParser(
		cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)))

	slog.Info("Starting cron job")

	// Задача, выполняемая раз в сутки в 2:30
	c.AddFunc("0 6 14 * * *", func() {
		slog.Info("Задача выполняется раз в день в 14:06")
	})

	// Задача, выполняемая каждый час
	c.AddFunc("@hourly", func() {
		slog.Info("Задача выполняется каждый час")
	})

	// Задача, выполняемая каждые 30 минут
	c.AddFunc("0 0,30 * * * *", func() {
		slog.Info("Задача выполняется каждые 30 минут")
	})

	// Задача, выполняемая каждую минуту
	c.AddFunc("0 * * * * *", func() {
		slog.Info("Задача выполняется каждую минуту")
	})

	// Задача, выполняемая каждую секунду
	c.AddFunc("* * * * * *", func() {
		slog.Info("Задача выполняется каждую секунду")
	})

	// Задача, выполняемая каждые 10 секунд
	c.AddFunc("*/10 * * * * *", func() {
		slog.Info("Задача выполняется каждые 10 секунд")
	})

	// Альтернативный способ для задачи каждые 30 минут
	// c.AddFunc("@every 30m", func() {
	//     fmt.Println("Задача выполняется каждые 30 минут")
	// })

	// Запускаем планировщик
	c.Start()

	for true {
		select {
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
