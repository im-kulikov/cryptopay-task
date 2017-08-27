package main

import (
	"bytes"
	"io/ioutil"
	"regexp"
)

// На всякий случай, нам стоит
var r = regexp.MustCompile(`[\s,.!?]+`)

func readWords(fPath string) (words map[string]int, err error) {
	// Читаем файл целиком:
	buf, err := ioutil.ReadFile(fPath)

	words = make(map[string]int)

	if err != nil {
		return
	}

	// Удаляем лишние пробелы:
	buf = bytes.TrimSpace(buf)
	// Возводим в верхний регистр:
	buf = bytes.ToUpper(buf)

	// Разбиваем по разделителям:
	for _, word := range r.Split(string(buf), -1) {
		if _, ok := words[word]; ok {
			words[word]++
		} else {
			words[word] = 1
		}
	}

	return
}
