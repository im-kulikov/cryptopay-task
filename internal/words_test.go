package internal

import (
	"fmt"
	"testing"
)

func BenchmarkReadWords(b *testing.B) {
	var (
		err   error
		words = make(Words)
	)

	fmt.Println("BenchmarkReadWords")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err = ReadFile("../data/187", func(word string) {
			if _, ok := words[word]; ok {
				words[word]++
			} else {
				words[word] = 1
			}
		}); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkReadWordsBuffer(b *testing.B) {
	var (
		err   error
		words = make(Words)
	)

	fmt.Println("BenchmarkReadWordsBuffer")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err = readFile("../data/187", func(word string) {
			if _, ok := words[word]; ok {
				words[word]++
			} else {
				words[word] = 1
			}
		}); err != nil {
			b.Fatal(err)
		}
	}
}
