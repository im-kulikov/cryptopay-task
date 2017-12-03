package main

import (
	"flag"
	"fmt"
	"os"
	_ "runtime/pprof"
	"time"

	"github.com/im-kulikov/cryptopay-task/internal"
	"github.com/pkg/profile"
)

const (
	//defaultFilename = "data/example_input"
	//defaultFilename   = "data/187"
	defaultFilename   = "data/187"
	defaultVocabulary = "data/vocabulary.txt"
)

var (
	debug   = flag.Bool("debug", false, "enable debug")
	file    = flag.String("file", defaultFilename, "file to process")
	vocFile = flag.String("vocabulary", defaultVocabulary, "vocabulary file")
)

func main() {
	flag.Parse()

	if *debug {
		defer profile.Start(
			profile.Quiet,
			profile.TraceProfile,
			//profile.MemProfile,
			//profile.MemProfileRate(2048),
			//profile.CPUProfile,
			profile.ProfilePath("./github"),
		).Stop()
	}

	var (
		err        error
		words      internal.Words
		vocabulary internal.Vocabulary
		notify     chan int
		distance   int
		index      int
		size       = 30
	)

	if _, err = os.Open(*file); err != nil || os.IsNotExist(err) {
		fmt.Printf("%q not exists", *file)
		os.Exit(0)
	}

	if _, err = os.Open(*vocFile); err != nil || os.IsNotExist(err) {
		fmt.Printf("%q not exists", *vocFile)
		os.Exit(0)
	}

	t := time.Now()

	if words, err = internal.ReadWords(*file); err != nil {
		panic(err)
	}

	if vocabulary, err = internal.ReadVocabulary(*vocFile, words); err != nil {
		panic(err)
	}

	length := len(words)

	if length < size {
		size = length
	}

	notify = make(chan int, size)

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

	fmt.Printf("Distance: %d\n", distance)
	fmt.Printf("Spent: %s\n", time.Now().Sub(t))
}
