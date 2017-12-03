package internal

import (
	"os"
	"strings"
	"text/scanner"
)

type Vocabulary map[int][]string

func ReadVocabulary(filename string, words Words) (vocabulary Vocabulary, err error) {
	var (
		file *os.File
		scan scanner.Scanner
	)

	if file, err = os.Open(filename); err != nil {
		return
	}

	scan.Init(file)
	vocabulary = make(Vocabulary)

	for tok := scan.Scan(); tok != scanner.EOF; tok = scan.Scan() {
		word := strings.ToUpper(scan.TokenText())
		if _, ok := words[word]; ok {
			delete(words, word)
			continue
		}
		length := len(word)
		vocabulary[length] = append(vocabulary[length], word)
	}

	return
}

func (v Vocabulary) partSearch(word string, wordLen int) int {
	distMin := wordLen

	for _, s := range v[wordLen] {
		if dist := levenshtein(word, s); dist == 1 {
			return 1
		} else if distMin > dist {
			distMin = dist
		}
	}

	return distMin
}

func (v Vocabulary) stepSearch(word string, step int, wordLen int) int {
	if _, ok := v[step]; !ok {
		return wordLen
	}

	return v.partSearch(word, step)
}
