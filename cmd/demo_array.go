package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5} // Slice
	//n2 := [5]int{1, 2, 3, 4, 5} // Array
	fmt.Println(n[0:3])
	n = append(n, 6)
	for i := 0; i < len(n); i++ {
		fmt.Println(n[i])
	}

	// for i, v := range n {
	// 	fmt.Println(i, v)
	// }

	// i := 0
	// for i < len(n) {
	// 	fmt.Println(n[i])
	// 	i++
	// }
}
