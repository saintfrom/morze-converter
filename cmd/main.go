package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/morse-converter/server"
)

func main() {
	// Создаём логгер
	logger := log.New(os.Stdout, "MORSE: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Создаём сервер
	srv := server.New(logger)

	// Запускаем сервер
	if err := srv.Start(); err != nil {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
