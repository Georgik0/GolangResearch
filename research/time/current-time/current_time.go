package current_time

import (
	"fmt"
	"time"
)

func Research() {
	// Устанавливаем время начала (например, Эпоха Unix - 1 января 1970 года)
	startTime := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)

	// Получаем текущее время
	currentTime := time.Now()

	// Вычисляем разницу между текущим временем и временем начала
	elapsed := currentTime.Sub(startTime)

	fmt.Printf("Absolute time elapsed since start: %v секунд %v миллисекунд\n", int(elapsed.Seconds()), int(elapsed.Nanoseconds()/1e6))

	customTime := time.Unix(1611847615, 705248000)
	fmt.Println(customTime.Unix())
	fmt.Println(time.Now().Unix())

	deltha := currentTime.Sub(customTime).Hours()

	fmt.Println(deltha / 24)
}
