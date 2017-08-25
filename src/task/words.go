package main

import (
	"bytes"
	"io/ioutil"
	"regexp"
)

// На всякий случай, нам стоит
var r = regexp.MustCompile(`[\s,.!?]+`)

func readWords(fPath string) (words []string, err error) {
	// Читаем файл целиком:
	buf, err := ioutil.ReadFile(fPath)

	if err != nil {
		return
	}

	// Удаляем лишние пробелы:
	buf = bytes.TrimSpace(buf)
	// Возводим в верхний регистр:
	buf = bytes.ToUpper(buf)

	// Разбиваем по разделителям:
	words = r.Split(string(buf), -1)

	return
}
