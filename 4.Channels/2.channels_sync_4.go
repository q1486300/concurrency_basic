package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 10000; i++ {
		wg.Add(1)
		go func(j int) {
			var result = 0

			goChan := make(chan int)
			mainChan := make(chan string)

			calculateSquare := func() {
				time.Sleep(time.Second * 3)
				result = j * j
				goChan <- result
			}
			reportResult := func() {
				fmt.Printf("The result of %d square is %d\n", j, <-goChan)
				// 阻塞，直到可以從 goChan 讀取值出來
				mainChan <- "You can quit now. I'm done."
			}

			go calculateSquare()
			go reportResult()
			<-mainChan // 阻塞，直到可以從 mainChan 讀取值出來
			wg.Done()
		}(i)
	}
	wg.Wait()
}
