package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	runtime.GOMAXPROCS(1) // 即使只有一個處理器，也可以併發執行 I/O bound 操作

	links := []string{
		"https://www.google.com/",
		"https://go.dev/",
		"https://medium.com/",
		"https://github.com/",
		"https://stackoverflow.com/",
		"https://www.mysql.com/",
		"https://gin-gonic.com/",
		"https://gorm.io/",
	}
	wg.Add(len(links))

	start := time.Now()

	for _, link := range links {
		go checkLink(link)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Processes took %v\n", elapsed)
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is not responding!")
		return
	}
	fmt.Println(link, "is LIVE!")
	wg.Done()
}
