package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func generateAPIKey(length int) (string, error) {
	// Вычисляем количество байт, необходимое для создания строки заданной длины.
	byteLength := length / 2

	// Создаем буфер для случайных байтов.
	randomBytes := make([]byte, byteLength)

	// Заполняем буфер случайными байтами.
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Преобразуем случайные байты в строку в шестнадцатеричном формате.
	apiKey := hex.EncodeToString(randomBytes)

	return apiKey, nil
}

func main() {
	// Укажите желаемую длину API-ключа в байтах (в этом примере - 32 байта).
	apiKeyLength := 32

	// Генерируем API-ключ.
	apiKey, err := generateAPIKey(apiKeyLength)
	if err != nil {
		fmt.Println("Ошибка при генерации API-ключа:", err)
		return
	}

	fmt.Println("Сгенерированный API-ключ:", apiKey)
}
