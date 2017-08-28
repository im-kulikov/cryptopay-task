package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	vWords      map[int][]string
	vWordsLen   int
	fWords      map[string]int
	fWordsByLen map[int][]string
	fWordsLen   int
	distance    int
	notify      chan int
)

func main() {
	t := time.Now()

	// Если не предоставили входного файла:
	if len(os.Args) < 2 {
		log.Fatal("need input file")
	}

	// Читаем слова:
	readWords()
	// Читаем словарь:
	readVocabulary()

	fWordsLen = len(fWords)

	// Канал размером <кол.слов> / 3:
	notify = make(chan int)

	for word, count := range fWords {
		go getDistanceFast(word, count)
	}

	j := 0

	for dist := range notify {
		distance += dist

		j++

		if j == fWordsLen {
			close(notify)
		}
	}

	println(distance)

	fmt.Printf("Spent: %s\n", time.Now().Sub(t))
}
