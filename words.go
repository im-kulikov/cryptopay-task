package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func readWords() {
	if len(os.Args) < 2 {
		log.Fatal("Error: no input file")
	}

	// Читаем файл целиком:
	buf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	fWords = make(map[string]int)

	// Возводим в верхний регистр:
	buf = bytes.ToUpper(buf)

	// Разбиваем строку по словам (https://golang.org/pkg/strings/#Fields):
	for _, word := range strings.Fields(string(buf)) {
		// Если слово уже есть - добавляем количество,
		// иначе единица:
		if _, ok := fWords[word]; ok {
			fWords[word]++
		} else {
			fWords[word] = 1
		}
	}
}
