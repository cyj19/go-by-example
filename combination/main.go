package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for index, v := range vs {
		if v == t {
			return index
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	result := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Map(vs []string, f func(string) string) []string {
	result := make([]string, len(vs))
	for index, v := range vs {
		result[index] = f(v)
	}
	return result
}

func main() {
	vs := []string{"peach", "apple", "pear", "plum"}

	fmt.Println(vs)

	fmt.Println(Index(vs, "pear"))

	fmt.Println(Include(vs, "dog"))

	fmt.Println(Any(vs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))

	fmt.Println(All(vs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))

	fmt.Println(Filter(vs, func(s string) bool {
		return strings.Contains(s, "e")
	}))

	fmt.Println(Map(vs, strings.ToUpper))
}
