package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

// VOCABULARY - путь к словарю:
const VOCABULARY = "data/vocabulary.txt"

func readVocabulary() {
	var (
		vPath string
	)

	// Получаем путь к словарю:
	vPath, _ = filepath.Abs(VOCABULARY)

	buf, err := ioutil.ReadFile(vPath)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	vWords = make(map[int][]string)

	for _, word := range strings.Fields(string(buf)) {
		wordLen := len(word)

		if _, ok := fWords[word]; ok {
			delete(fWords, word)
		}

		vWords[wordLen] = append(vWords[wordLen], word)
	}

	return
}
