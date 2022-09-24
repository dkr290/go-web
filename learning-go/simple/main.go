package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
			}
			countLines(f, counts)
			f.Close()
		}

		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%v\n", n, line)
			}
		}
	}

}

func countLines(f *os.File, c map[string]int) {
	inp := bufio.NewScanner(f)
	for inp.Scan() {
		c[inp.Text()]++

	}
}
