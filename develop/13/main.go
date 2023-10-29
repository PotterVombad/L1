package main


func main() {
	a := 1
	b := 10
	println(a, b)
	a, b = b, a
	println(a, b)
}