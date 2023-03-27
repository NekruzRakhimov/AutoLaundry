package utils

import (
	"errors"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func FuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func StringToArray(s string) ([]int, error) {
	s = strings.Trim(s, "{}")      // удаление фигурных скобок вокруг массива
	parts := strings.Split(s, ",") // разделение строки по запятой
	result := make([]int, len(parts))
	for i, p := range parts {
		val, err := strconv.Atoi(strings.TrimSpace(p)) // конвертация строки в инт
		if err != nil {
			return nil, err
		}
		result[i] = val
	}
	return result, nil
}

func ValidateDateFormat(dateString string) error {
	// Проверяем, соответствует ли строка формату "YYYY-MM-DD"
	if matched, _ := regexp.MatchString(`\d{4}-\d{2}-\d{2}`, dateString); !matched {
		return errors.New("Invalid date format")
	}

	// Парсим строку в объект time.Time и проверяем, что это действительно дата
	if _, err := time.Parse("2006-01-02", dateString); err != nil {
		return err
	}

	return nil
}
