package main

import (
	"os"
	"runtime"
)

var (
	vWords    []string
	vWordsLen int
	fWords    []string
	fWordsLen int
)

func main() {
	var (
		err error
	)

	//t := time.Now()

	// Если не предоставили входного файла:
	if len(os.Args) < 2 {
		println("need input file")
		os.Exit(0)
	}

	// Если входного файла не существует:
	if _, err = os.Stat(os.Args[1]); os.IsNotExist(err) {
		println("Error: input file not exists")
		os.Exit(0)
	}

	// Читаем слова из входного файла:
	if fWords, err = readWords(os.Args[1]); err != nil {
		println("Error: ", err)
		os.Exit(0)
	}

	fWordsLen = len(fWords)

	// Если входной файл пуст:
	if fWordsLen == 0 {
		println("Error: the input file does not contain the words")
		os.Exit(0)
	}

	// Читаем словарь:
	if vWords, err = newVocabulary(); err != nil {
		println("Error: ", err)
		os.Exit(0)
	}

	vWordsLen = len(vWords)

	// Если словарь пуст:
	if vWordsLen == 0 {
		println("Error: the vocabulary does not contain the words")
		os.Exit(0)
	}

	distance := 0

	notify := make(chan int)

	for i := 0; i < fWordsLen; i++ {
		go func(j int) {
			notify <- calculateDistance(j)
		}(i)

		runtime.Gosched()
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

	//fmt.Printf("Spent: %s\n", time.Now().Sub(t))
}
