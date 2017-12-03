package internal

import (
	"os"
	"strings"
	"text/scanner"
)

type Words = map[string]int

func ReadWords(filename string) (words Words, err error) {
	var (
		file *os.File
		scan scanner.Scanner
	)

	if file, err = os.Open(filename); err != nil {
		return
	}

	scan.Init(file)
	words = make(Words)

	for tok := scan.Scan(); tok != scanner.EOF; tok = scan.Scan() {
		//fmt.Printf("%s: %s\n", scan.Position, scan.TokenText())
		word := strings.ToUpper(scan.TokenText())
		if _, ok := words[word]; ok {
			words[word]++
		} else {
			words[word] = 1
		}
	}

	return
}
