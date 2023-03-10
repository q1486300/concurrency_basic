package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(8) // 8個處理器，同時處理 8 個 goroutines

	c := make(chan string)

	start := time.Now()
	go counta(c)
	go countb(c)
	go countc(c)
	go countd(c)
	go counte(c)
	go countf(c)
	go countg(c)
	go counth(c)

	for i := 0; i < 8; i++ {
		fmt.Println(<-c)
	}

	elapsed := time.Since(start)
	fmt.Printf("Processes took %v\n", elapsed)
}

func counta(c chan string) {
	fmt.Println("AAAA is starting")
	for i := 0; i < 10_000_000_000; i++ {

	}
	c <- "AAAA is done"
}

func countb(c chan string) {
	fmt.Println("BBBB is starting")
	for i := 0; i < 10_000_000_000; i++ {

	}
	c <- "BBBB is done"
}

func countc(c chan string) {
	fmt.Println("CCCC is starting")
	for i := 0; i < 10_000_000_000; i++ {

	}
	c <- "CCCC is done"
}

func countd(c chan string) {
	fmt.Println("DDDD is starting")
	for i := 0; i < 10_000_000_000; i++ {

	}
	c <- "DDDD is done"
}

func counte(c chan string) {
	fmt.Println("EEEE is starting")
	for i := 0; i < 10_000_000_000; i++ {

	}
	c <- "EEEE is done"
}

func countf(c chan string) {
	fmt.Println("FFFF is starting")
	for i := 0; i < 10_000_000_000; i++ {

	}
	c <- "FFFF is done"
}

func countg(c chan string) {
	fmt.Println("GGGG is starting")
	for i := 0; i < 10_000_000_000; i++ {

	}
	c <- "GGGG is done"
}

func counth(c chan string) {
	fmt.Println("HHHH is starting")
	for i := 0; i < 10_000_000_000; i++ {

	}
	c <- "HHHH is done"
}
