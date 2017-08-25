package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"path/filepath"
)

const VOCABULARY = "/../data/vocabulary.txt"

func newVocabulary() (vWords []string, err error) {
	var (
		vPath    string
		basePath string
		file     *os.File
		part     []byte
		prefix   bool
	)

	// Получаем базовый путь:
	basePath, err = filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		return
	}

	// Получаем путь к словарю:
	vPath, err = filepath.Abs(basePath + VOCABULARY)

	if err != nil {
		return
	}

	// Проверяем наличие файла:
	if file, err = os.Open(vPath); err != nil {
		return
	}

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 1024))

	// Читаем файл кусками:
	for {

		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}

		buffer.Write(part)

		if !prefix {
			vWords = append(vWords, buffer.String())
			buffer.Reset()
		}

	}

	// Игнорируем EOF:
	if err == io.EOF {
		err = nil
	}

	return
}
