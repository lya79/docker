package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

}

func channelClose() {
	ch := make(chan int)

	go func(ch chan int) {
		ch <- 1
		close(ch)
	}(ch)

	for {
		v, ok := <-ch
		if !ok {
			log.Println("break")
			break
		}
		log.Println("v:", v)
	}
}

func selectRormal() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func selectRange() {
	fibonacci := func(n int, c chan int) {
		x, y := 1, 1
		for i := 0; i < n; i++ {
			c <- x
			x, y = y, x+y
		}
		close(c)
	}

	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
func selectTimeout() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}

func channel() {
	ch := make(chan int) // buffer 0, 如果寫入了,  沒人拿走就會一直卡在寫入
	ch <- 1              // 會卡在這行
	fmt.Println("hello") //在此行执行之前Go就会报死锁

	ch2 := make(chan int, 2) // buffer 2, 如果寫入了,  允許寫兩個並且沒人拿走, 如果寫第三個時還有兩個沒拿走就會卡住
	ch2 <- 1
	ch2 <- 1
	ch2 <- 1             // 會卡在這行
	fmt.Println("hello") //在此行执行之前Go就会报死锁
}
