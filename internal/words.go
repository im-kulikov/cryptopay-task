package internal

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

type Words = map[string]int

type Handler func(word string)

func readFile(filename string, handler Handler) (err error) {
	var (
		file *os.File
		scan *bufio.Scanner
	)
	if file, err = os.Open(filename); err != nil {
		return
	}
	scan = bufio.NewScanner(file)
	scan.Split(bufio.ScanWords)
	//buf := make([]byte, 0, 64*1024)
	//scan.Buffer(buf, 1024*1024)
	for scan.Scan() {
		handler(scan.Text())
		//parts := strings.Fields(line)
		//
		//for i := range parts {
		//	handler(parts[i])
		//}
	}
	err = scan.Err()
	return
}

func ReadFile(filename string, handler Handler) (err error) {
	var data []byte
	if data, err = ioutil.ReadFile(filename); err != nil {
		return
	}
	for _, word := range strings.Fields(string(data)) {
		handler(word)
	}
	return
}
