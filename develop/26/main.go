package main

import (
	"fmt"
	"strings"
)

func Unique(s string) bool{
	lowString := strings.ToLower(s)
	uniq := make(map[rune]bool)
	for _, v := range lowString{
		_, ok := uniq[v]
		if !ok{
			uniq[v] = true
			continue
		}
		return false
	}
	return true
}
func main(){
	fmt.Println(Unique("abcd"))
	fmt.Println(Unique("abCdefAaf"))
	fmt.Println(Unique("aabcd"))

}