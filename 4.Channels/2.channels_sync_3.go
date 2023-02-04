package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		var result = 0
		var value = 97

		goChan := make(chan int)
		mainChan := make(chan string)

		calculateSquare := func() {
			time.Sleep(time.Second * 3)
			result = value * value
			goChan <- result
		}
		reportResult := func() {
			fmt.Printf("The result of %d square is %d\n", value, <-goChan)
			// 阻塞，直到可以從 goChan 讀取值出來
			mainChan <- "You can quit now. I'm done."
		}

		go calculateSquare()
		go reportResult()
		<-mainChan // 阻塞，直到可以從 mainChan 讀取值出來
		wg.Done()
	}()
	wg.Wait()
}
