package main

import (
	"errors"
	"fmt"
)

func main() {
	r, err := doSth()
	fmt.Println(r, err)
}

var myError = errors.New("My new Error")

type MyError struct {
	code   int
	detail string
}

func (me MyError) Error() string {
	return fmt.Sprintf("My error with code %v and detail %v", me.code, me.detail)
}

func doSth() (string, error) {
	result := "result"
	err := MyError{code: 404, detail: "Not found"}
	return result, err
}
