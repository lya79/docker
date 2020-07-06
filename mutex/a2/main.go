package main

import (
	"log"
	"time"
)

func main(){
	ch := make(chan int)

	// input
	go func(ch chan int){
		for i:=0;i<10;i++{
			ch <- 1
			log.Println("input")
		}
	}(ch)

	// output
	go func(ch chan int){
		for {
			_, ok := <- ch
			if !ok{
				return
			}
			log.Println("output")
		}
	}(ch)


	time.Sleep(1*time.Second)
}