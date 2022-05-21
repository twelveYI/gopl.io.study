package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	args := os.Args
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}
