package main

import "fmt"

type Set struct {
	el []string
}

func MakeSet(s []string) Set {
	dict := make(map[string]struct{})
	for _, v := range s {
		_, ok := dict[v]
		if !ok {
			dict[v] = struct{}{}
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
