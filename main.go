package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var lines, bytes bool
	flag.BoolVar(&lines, "l", false, "Count lines")
	flag.BoolVar(&bytes, "b", false, "Count bytes")
	flag.Parse()

	counter, err := count(os.Stdin, lines, bytes)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(counter)
}

func count(r io.Reader, lines, bytes bool) (int, error) {
	scanner := bufio.NewScanner(r)

	if !lines && !bytes {
		scanner.Split(bufio.ScanWords)
	}

	if bytes {
		scanner.Split(bufio.ScanBytes)
	}

	counter := 0
	for scanner.Scan() {
		counter++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return counter, nil
}
