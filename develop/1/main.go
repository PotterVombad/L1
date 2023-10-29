package main

import "fmt"

type Human struct{
	name string
}

func (h Human) GetName() {
	fmt.Println(h.name)
}

type Action struct {
	Human
}


func main() {
	a := Human{"Sergey"}
	b := Action{a}
	b.GetName()
}