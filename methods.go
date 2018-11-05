package main

import "fmt"

type A1 struct {
	B1
	//Name string
}
type B1 struct {
	Name string
}

func main() {
	//a := A1{Name: "A1", B1: B1{Name: "B1"}}
	a := A1{B1: B1{Name: "B1"}}
	fmt.Println(a.Name, a.B1.Name)
	a1 := A1{}
	a1.Print()
	b1 := B1{}
	b1.Print()
}

func (a1 A1) Print() {
	fmt.Println("A1")
}

func (b1 B1) Print() {
	fmt.Println("B1")
}
