package internal

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func BenchmarkLevenshtein(b *testing.B) {
	var (
		err   error
		words = make(Words)
	)

	fmt.Println("BenchmarkLevenshtein")

	if err = ReadFile("../data/187", func(word string) {
		word = strings.ToUpper(word)
		if len(word) == 1 {
			return
		}

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
		for word := range words {
			tested := []rune(word)

			l := rand.Intn(len(word) - 1)
			tested[l] = ([]rune("A"))[0]

			if val := levenshteinDistance(word, string(tested)); val != 1 && word != string(tested) {
				b.Fatal(val, string(tested), word)
			}
		}
	}
}
