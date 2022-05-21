package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var counts = make(map[string]int)
	if 0 == len(os.Args[1:]) {
		// 从标准输入中读取
		countLines(counts, os.Stdin)
	} else {
		// 从文件中读取，文明名称为命令参数
		for _, file := range os.Args[1:] {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(counts, f)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}

	}
}

func countLines(counts map[string]int, file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		key := scanner.Text()
		counts[key]++
	}
}
