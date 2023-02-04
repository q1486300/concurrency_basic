package main

import (
	"fmt"
	"time"
)

var result = 0
var value = 97

func main() {
	goChan := make(chan int)
	mainChan := make(chan string)
	go calculateSquare(value, goChan)
	go reportResult(goChan, mainChan)
	<-mainChan // 阻塞，直到可以從 mainChan 讀取值出來
}

func calculateSquare(value int, goChan chan int) {
	fmt.Println("Calculating for 3 seconds...")
	time.Sleep(time.Second * 3)
	result = value * value
	goChan <- result
}

func reportResult(goChan chan int, mainChan chan string) {
	time.Sleep(time.Second)
	fmt.Printf("The result of %d square is %d\n", value, <-goChan)
	// 阻塞，直到可以從 goChan 讀取值出來
	mainChan <- "You can quit now. I'm done."
}
