package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var lines bool
	flag.BoolVar(&lines, "l", false, "Count lines")
	flag.Parse()

	counter, err := count(os.Stdin, lines)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(counter)
}

func count(r io.Reader, lines bool) (int, error) {
	scanner := bufio.NewScanner(r)

	if !lines {
		scanner.Split(bufio.ScanWords)
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
