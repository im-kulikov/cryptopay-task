package main

import (
	"sort"

	"fuzzy"
)

func calculateDistance(i int) int {
	wordSrc := fWords[i]
	distMin := len(wordSrc)
	wordLen := len(wordSrc)

	matches := fuzzy.RankFind(wordSrc, vWords)

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

			if dist := fuzzy.LevenshteinDistance(wordSrc, wordDst); dist < distMin {
				distMin = dist
			}

			if distMin == 0 {
				break
			}
		}
	}

	return distMin
}
