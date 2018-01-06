package internal

import (
	"fmt"
	"testing"
)

func BenchmarkReadVocabulary(b *testing.B) {
	var (
		err        error
		words      = make(Words, 0)
		vocabulary = make(Vocabulary, 0)
	)
	fmt.Println("BenchmarkReadVocabulary")

	if err = ReadFile("../data/187", func(word string) {
		if _, ok := words[word]; ok {
			words[word]++
		} else {
			words[word] = 1
		}
	}); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ReadFile("../data/vocabulary.txt", func(word string) {
			if _, ok := words[word]; ok {
				delete(words, word)
			}
			length := len(word)
			vocabulary[length] = append(vocabulary[length], word)
		})
	}
}

func BenchmarkReadVocabularyBuffered(b *testing.B) {
	var (
		err        error
		words      = make(Words, 0)
		vocabulary = make(Vocabulary, 0)
	)
	fmt.Println("BenchmarkReadVocabularyBuffered")

	if err = readFile("../data/187", func(word string) {
		if _, ok := words[word]; ok {
			words[word]++
		} else {
			words[word] = 1
		}
	}); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		readFile("../data/vocabulary.txt", func(word string) {
			if _, ok := words[word]; ok {
				delete(words, word)
			}
			length := len(word)
			vocabulary[length] = append(vocabulary[length], word)
		})
	}
}
