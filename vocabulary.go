package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

// VOCABULARY - путь к словарю:
const VOCABULARY = "data/vocabulary.txt"

func readVocabulary() {
	var (
		vPath string
	)

	// Получаем путь к словарю:
	vPath, _ = filepath.Abs(VOCABULARY)

	//buf, err := ioutil.ReadFile(vPath)

	buf, err := os.Open(vPath)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	vWords = make(map[int][]string)

	scanner := bufio.NewScanner(buf)

	for scanner.Scan() {
		word := scanner.Text()
		wordLen := len(word)

		if _, ok := fWords[word]; ok {
			delete(fWords, word)
		}

		vWords[wordLen] = append(vWords[wordLen], word)
	}

	return
}
