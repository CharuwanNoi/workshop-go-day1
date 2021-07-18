package main

import (
	"day01"
	"fmt"
)

type parent struct {
}

type child struct {
	parent
}

func (p parent) callParent() {
	fmt.Println("call Parent...")
}

// func (c child) callParent() {
// 	c.parent.callParent() // Call from parent
// 	fmt.Println("call parent from child...")
// }

func main() {
	c := child{}
	c.callParent()

	// d := day01.Department{"demo"}
	// p := day01.Person{Id: 1, Name: "charuwan", Department: d}
	p := day01.NewPerson(1, "charuwan", "demo")

	p.DoStm()
	fmt.Println(p.DevName)
	fmt.Println(p.Department.DevName)
	//fmt.Printf("%+v", p)
}
