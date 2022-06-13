package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		if !seen[s] {
			seen[s] = true
			fmt.Println(s)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup:%v\n", err)
		os.Exit(1)
	}
}
