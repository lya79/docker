package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var capacity int // 倉庫最大容量
	var num int      // 商品數量

	capacity = 3
	num = 0

	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	go func() {
		for {
			mutex.Lock()

			for {
				if num >= capacity {
					fmt.Println("倉庫放不下", "目前有", num, "個")
					cond.Wait()
					continue
				}
				break
			}

			num = num + 1
			fmt.Println("增加一個", "目前有", num, "個")

			// cond.Signal()
			cond.Broadcast()

			mutex.Unlock()
		}
	}()

	go func() {
		for {
			mutex.Lock()

			for {
				if num <= 0 {
					fmt.Println("倉庫清空了", "目前有", num, "個")
					cond.Wait()
					continue
				}
				break
			}

			num = num - 1
			fmt.Println("拿走一個", "目前有", num, "個")

			// cond.Signal()
			cond.Broadcast()

			mutex.Unlock()
		}
	}()

	time.Sleep(1 * time.Second)
}
