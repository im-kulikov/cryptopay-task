package internal

import "testing"

func BenchmarkReadWordsReadWords(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ReadWords("../data/187")
	}
}
