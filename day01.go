package day01

import "fmt"

func dosometing(input int) func(i int) int {
	return func(x int) int { return input * 2 }
}

func Hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
