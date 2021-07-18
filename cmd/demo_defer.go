package main

import "fmt"

func main() {
	fmt.Println("Start")
	defer fmt.Println("TODO")
	fmt.Println("Finish")
}
