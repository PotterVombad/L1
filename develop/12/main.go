package main

import "fmt"

type Set struct {
	el []string
}

func MakeSet(s []string) Set {
	dict := make(map[string]bool)
	for _, v := range s {
		_, ok := dict[v]
		if !ok {
			dict[v] = true
		}
	}
	keys := make([]string, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}
	return Set {
		el: keys,
	}
}

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(MakeSet(str))
}
