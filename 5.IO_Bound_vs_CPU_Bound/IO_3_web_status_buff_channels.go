package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

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
	c := make(chan string, len(links))

	start := time.Now()

	for _, link := range links {
		go checkLink(link, c)
	}

	for len(c) < len(links) {

	}
	//for range links {
	//	fmt.Println("channel message:", <-c)
	//}

	elapsed := time.Since(start)
	fmt.Printf("Processes took %v\n", elapsed)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is not responding!")
		c <- fmt.Sprintf("%s is not responding!", link)
		return
	}
	fmt.Println(link, "is LIVE!")
	c <- fmt.Sprintf("%s is LIVE!", link)
}
