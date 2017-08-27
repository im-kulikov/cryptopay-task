package main

import (
	"sort"

	"fuzzy"
)

func calculateDistance(word string) int {
	distMin := len(word)
	wordLen := len(word)

	matches := fuzzy.RankFind(word, vWords)

	if len(matches) > 0 {
		sort.Sort(matches)

		distMin = matches[0].Distance
	}

	if distMin > 1 {
		for j := 0; j < vWordsLen; j++ {
			wordDst := vWords[j]

			// Для больших входящик файлов:
			dstLen := len(wordDst)
			if dstLen < wordLen-4 || dstLen > wordLen+4 {
				continue
			}

			if wordDst == word {
				distMin = 0
				break
			}

			if dist := fuzzy.LevenshteinDistance(word, wordDst); dist < distMin {
				distMin = dist
			}

			if distMin == 0 {
				break
			}
		}
	}

	return distMin
}
