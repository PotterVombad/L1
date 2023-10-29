package main

import (
	"fmt"
	"math/big"
)

func Sum(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func Multiply(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func Subtract(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func Divide(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func main() {
	a := "136541365136543521651321165132165063521035105"
    aBig := new(big.Int)
    aBig.SetString(a, 10)
	b := "92233720368547758079223372036854775807" 
	bBig := new(big.Int)
    bBig.SetString(b, 10)

	fmt.Println(Sum(aBig, bBig))
	fmt.Println(Subtract(aBig, bBig))
	fmt.Println(Multiply(aBig, bBig))
	fmt.Println(Divide(aBig, bBig))
}