package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, v := range vals {
		total += v
	}
	return total
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2, 3))

	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))
}
