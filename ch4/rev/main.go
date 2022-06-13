package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	input := bufio.NewScanner(os.Stdin)

outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))

		}
		reverse(ints)
		fmt.Printf("%v\n", ints)

	}
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
