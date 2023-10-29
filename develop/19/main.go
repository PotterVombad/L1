package main

import "fmt"

func ReverseRune(s string) string {
	answer := []rune(s)
    for first, last := 0, len(answer)-1; first < last; first, last = first+1, last-1 {
        answer[first], answer[last] = answer[last], answer[first]
    }
    return string(answer)
}

func main() {
    example := "главрыба"
    fmt.Println(ReverseRune(example))
}