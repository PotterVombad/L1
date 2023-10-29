//Какой размер у структуры struct{}{}?
//Ответ: 0
package main

import (
    "fmt"
    "unsafe"
)


func main() {
    s2 := struct{}{}
    var s1 interface{}

    fmt.Printf("s2 size: %v\n", unsafe.Sizeof(s2))
    fmt.Printf("s2 size: %v\n", unsafe.Sizeof(s1))
}