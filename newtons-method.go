package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x/2.0
	for i:=0; i<10; i++ {
		z0 := z
		z -= (z*z - x) / (2*z)
		fmt.Println(i)
		if z0== z {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(607))
}
