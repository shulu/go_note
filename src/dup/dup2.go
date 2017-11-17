// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	fmt.Println(len(files))
	if len(files) == 0 {
		countLines(os.Stdin, counts, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, files string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		var isCheck = counts[input.Text()]
		//fmt.Printf("%d\t%s\n", isCheck, input.Text())
		if isCheck > 0 {
			fmt.Println(files)
		}
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}


