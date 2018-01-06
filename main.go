package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/im-kulikov/cryptopay-task/internal"
	"github.com/pkg/profile"
)

const (
	//defaultFilename   = "data/187"
	defaultFilename   = "data/example_input"
	defaultVocabulary = "data/vocabulary.txt"
)

var (
	debug   = flag.Bool("debug", false, "enable debug")
	file    = flag.String("file", defaultFilename, "file to process")
	vocFile = flag.String("vocabulary", defaultVocabulary, "vocabulary file")
)

func spent(dt time.Time) {
	fmt.Println("Spent: ", time.Since(dt))
}

func main() {
	var (
		err        error
		notify     chan int
		distance   int
		index      int
		size       = 30
		words      = make(internal.Words)
		vocabulary = make(internal.Vocabulary)
	)

	flag.Parse()

	if *debug {
		defer profile.Start(
			profile.Quiet,
			//profile.TraceProfile,
			//profile.MemProfile,
			//profile.MemProfileRate(2048),
			profile.CPUProfile,
			profile.ProfilePath("./github"),
		).Stop()
	}

	defer spent(time.Now())

	defer func() {
		fmt.Println("Distance:", distance)
	}()

	if _, err = os.Open(*file); err != nil || os.IsNotExist(err) {
		fmt.Printf("%q not exists", *file)
		os.Exit(0)
	}

	if _, err = os.Open(*vocFile); err != nil || os.IsNotExist(err) {
		fmt.Printf("%q not exists", *vocFile)
		os.Exit(0)
	}

	if err = internal.ReadFile(*file, func(word string) {
		word = strings.ToUpper(word)
		if _, ok := words[word]; ok {
			words[word]++
		} else {
			words[word] = 1
		}
	}); err != nil {
		panic(err)
	}

	if err = internal.ReadFile(*vocFile, func(word string) {
		word = strings.ToUpper(word)
		if _, ok := words[word]; ok {
			delete(words, word)
		}
		length := len(word)
		vocabulary[length] = append(vocabulary[length], word)
	}); err != nil {
		panic(err)
	}

	length := len(words)

	if length < size {
		size = length
	}

	notify = make(chan int, size/3)

	// internal/levenstein: 295.641957ms
	// ------------------------------------
	// github.com/m1ome/leven: 307.011865ms
	//internal.SetLevenshtein(leven.Distance)
	// ------------------------------------
	// github.com/arbovm/levenshtein: 311.315178ms
	//internal.SetLevenshtein(levenshtein.Distance)
	// ------------------------------------
	// github.com/agext/levenshtein: 335.033736ms
	//internal.SetLevenshtein(func(a, b string) int {
	//	return levenshtein.Distance(a, b, levenshtein.NewParams())
	//})

	for word, count := range words {
		go vocabulary.Distance(word, count, notify)
	}

	for current := range notify {
		index++

		//fmt.Printf("Distance: %d : %d\n", distance, current)
		distance += current
		if index == length {
			close(notify)
		}
	}
}
