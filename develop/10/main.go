package main

import "fmt"

func TemperatureGroup(temp []float64) map[int][]float64 {
	answer := make(map[int][]float64)
	for _, v := range temp {
		key := (int(v) / 10) * 10
		answer[key] = append(answer[key], v)
	}
	return answer
}

func main() {
	str := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(TemperatureGroup(str))
}
