package dp

import (
	"bufio"
	"fmt"
	"os"
)

type line map[string]int

func dp() {
	counts := make(map[string]line)
	lines := make(line)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, lines)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, lines)
			counts[arg] = lines
			f.Close()
		}
	}
	for file, lines := range counts {
		for line, n := range lines {
			//for line, n := range content {
			//fmt.Println("line:", line, "n:", n)
			if n > 1 {
				fmt.Printf("%d\t%s\t%s\n", n, line, file)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
