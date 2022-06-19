package main

import (
	"fmt"
	"sort"
)

// 课程和先决课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, course := range items {
			if !seen[course] {
				seen[course] = true
				visitAll(m[course])
				order = append(order, course)
			}
		}
	}
	var keys []string
	// for k, _ := range m {
	// keys = append(keys, k)
	// }
	for k := range m {
		keys = append(keys, k)

	}
	sort.Strings(keys)
	visitAll(keys)

	return order
}
