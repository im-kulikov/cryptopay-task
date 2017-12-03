package internal

import "testing"

func BenchmarkReadVocabulary(b *testing.B) {
	var words, err = ReadWords("../data/187")

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ReadVocabulary("../data/vocabulary.txt", words)
	}
}
