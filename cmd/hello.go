package main

import (
	"day01"
	"day01/db"
	"fmt"
)

// type Day int // alias type

func main() {

	// var d Day
	// d = 1

	result := day01.Hello("charuwan")
	fmt.Println(result)

	row, err := db.SaveData("error")
	if err != nil {
		fmt.Println(row, err)
	}
}
