package main

import "fmt"

type Stringer interface {
	String() string
}

type processor interface {
	doSth()
}

type demo struct {
	id   int
	name string
}

func (d demo) doSth() {
	fmt.Println("Called do Sth")
}

func (d demo) String() string {
	return fmt.Sprintf("id=%v , name=%v", d.id, d.name)
}

func main() {
	//fmt.Printf("%+v", demo{})
	do(demo{})
	do(A(1))
}

type A int

func (a A) doSth() {
	fmt.Println("Called Do Sth from A")
}

func do(d processor) {
	d.doSth()
}
