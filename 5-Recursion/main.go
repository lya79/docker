package main

import "log"

func main() {
	log.Println(add(3)) // output: 6
}

// 累加遞迴
func add(value int) int {
	if value <= 1 {
		return 1
	}
	return value + add(value-1)
}
