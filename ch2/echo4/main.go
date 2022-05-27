package main

import (
	"flag"
	"fmt"
	"strings"
)

var n *bool = flag.Bool("n", true, "omit new line")
var sep *string = flag.String("s", ";", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
