package main

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
	answer := strings.Split(s, " ")
    for first, last := 0, len(answer)-1; first < last; first, last = first+1, last-1 {
        answer[first], answer[last] = answer[last], answer[first]
    }
    return strings.Join(answer, " ")
}

func main() {
    example := "snow dog sun"
    fmt.Println(Reverse(example))
}