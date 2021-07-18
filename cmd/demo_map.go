package main

import "fmt"

func main() {

	n := 1

	switch n {
	case 1:
		fmt.Println("OK 1")
	case 6:
		fmt.Println("OK 2")
	case 7:
		fmt.Println("OK 3")
	default:
		fmt.Println("default")
	}

	switch {
	case n == 1, n == 2:
		fmt.Println("OK 1")
	case n%2 == 0:
		fmt.Println("OK 2")
	case n > 10:
		fmt.Println("OK 3")
	default:
		fmt.Println("default")
	}

	grades := make(map[int]string)
	grades[100] = "A"
	grades[101] = "B"
	grades[102] = "C+"

	for k, v := range grades {
		fmt.Println(k, v)
	}

	fmt.Printf("Data fo 100=%v\n", grades[100])

	if v, found := grades[10011]; found {
		fmt.Println("Found =", v)
	} else {
		fmt.Println("Not Found")
	}
}
