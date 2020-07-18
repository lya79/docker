package main

import "fmt"

/*
- panic只会作用在所属的 goroutine
- 如果在某个 func内触发 panic,  触发 panic不管 func有没有 recover, 都会把剩下的 defer都执行完毕
- func會先執行 return賦值才跑 defer
*/
func deferCall() {
	defer func() {
		fmt.Println("B")
	}()
	defer func() {
		fmt.Println("C")
	}()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover from r:", r)
		}
	}()
	defer func() {
		fmt.Println("D")
	}()
	fmt.Println("E")
	panic("Panic")
	fmt.Println("F")
}

func deferStart() {
	deferCall()
	fmt.Println("A")
}

func main() {
	deferStart()
}
