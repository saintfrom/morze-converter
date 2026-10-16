package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/morse-converter/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		log.Printf("ошибка парсинга формы: %v", err)
		http.Error(w, "Ошибка парсинга формы", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		log.Printf("ошибка получения файла: %v", err)
		http.Error(w, "Ошибка получения файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("ошибка чтения файла: %v", err)
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	result, err := service.DetectAndConvert(string(data))
	if err != nil {
		log.Printf("ошибка конвертации: %v", err)
		http.Error(w, "Ошибка конвертации", http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("result_%s%s",
		time.Now().UTC().Format("20060102_150405.000000000"),
		ext,
	)

	out, err := os.Create(filename)
	if err != nil {
		log.Printf("ошибка создания файла: %v", err)
		http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	if _, err := out.WriteString(result); err != nil {
		log.Printf("ошибка записи в файл: %v", err)
		http.Error(w, "Ошибка записи в файл", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Конвертация выполнена успешно!\nФайл сохранён: %s\n\nРезультат:\n%s",
		filename, result)
}
