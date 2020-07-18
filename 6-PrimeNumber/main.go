package main

import (
	"fmt"
	"math"
)

func main() {
	prime(100)
}

func prime(value int) {
	for i := 2; i <= value; i++ {
		prime := true
		max := int(math.Sqrt(float64(i)))
		for k := 2; k <= max; k++ {
			if i%k == 0 {
				prime = false
				break
			}
		}
		if prime {
			fmt.Print(i, ", ")
		}
	}
}
