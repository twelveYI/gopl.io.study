package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) != 1 {
		for _, file := range files {
			bites, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
				continue
			}
			for _, line := range strings.Split(string(bites), "\n") {
				counts[line]++
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d,%s\n", n, line)
		}
	}
}
