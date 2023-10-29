package main

import (
	"fmt"
	"reflect"
)
//with switch - case
func FindType(smth interface{}) string {
	switch smth.(type) {
		case string: return "string"
		case int: return "int"
		case bool: return "bool"
		case chan int: return "chan int"
		case chan string: return "chan int"
		default: return "unknown"
	}
}
//with reflect lib
func FindTypeLib(smth interface{}) string {
	return reflect.TypeOf(smth).String()
}

func main() {
	str := "abcdef"
	numb := 43
	var chanel chan int
	bb := true
	fmt.Println(FindTypeLib(str), FindType(str))
	fmt.Println(FindTypeLib(numb), FindType(numb))
	fmt.Println(FindTypeLib(chanel), FindType(chanel))
	fmt.Println(FindTypeLib(bb), FindType(bb))
}