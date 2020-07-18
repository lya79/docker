package main

import "log"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	startIndex := 2
	endIndex := 4
	log.Println(arr[startIndex : endIndex+1]) // 3,4,5, // index從多少開始 : index之前

	//-----------------------
	// 複製問題

	// 有問題的方式
	a := []int{1, 2, 3}
	b := a
	a[0] = 99
	log.Println(b) // output: 99, 2, 3

	// 正確方式
	a = []int{1, 2, 3}
	copy(b, a)
	a[0] = 99
	log.Println(b) // output: 1, 2, 3
}
