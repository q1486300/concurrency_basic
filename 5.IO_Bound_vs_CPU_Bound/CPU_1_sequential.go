package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(1)
	//runtime.GOMAXPROCS(8)	// 串行程式，增加處理器個數不會提高效能

	start := time.Now()
	counta()
	countb()
	countc()
	countd()
	counte()
	countf()
	countg()
	counth()

	elapsed := time.Since(start)
	fmt.Printf("Processes took %v\n", elapsed)
}

func counta() {
	fmt.Println("AAAA is starting")
	for i := 0; i < 10_000_000_000; i++ {

	}
	fmt.Println("AAAA is done")
}

func countb() {
	fmt.Println("BBBB is starting")
	for i := 0; i < 10_000_000_000; i++ {

	}
	fmt.Println("BBBB is done")
}

func countc() {
	fmt.Println("CCCC is starting")
	for i := 0; i < 10_000_000_000; i++ {

	}
	fmt.Println("CCCC is done")
}

func countd() {
	fmt.Println("DDDD is starting")
	for i := 0; i < 10_000_000_000; i++ {

	}
	fmt.Println("DDDD is done")
}

func counte() {
	fmt.Println("EEEE is starting")
	for i := 0; i < 10_000_000_000; i++ {

	}
	fmt.Println("EEEE is done")
}

func countf() {
	fmt.Println("FFFF is starting")
	for i := 0; i < 10_000_000_000; i++ {

	}
	fmt.Println("FFFF is done")
}

func countg() {
	fmt.Println("GGGG is starting")
	for i := 0; i < 10_000_000_000; i++ {

	}
	fmt.Println("GGGG is done")
}

func counth() {
	fmt.Println("HHHH is starting")
	for i := 0; i < 10_000_000_000; i++ {

	}
	fmt.Println("HHHH is done")
}
