package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1) // 串行程式，增加處理器個數不會提高效能
	fmt.Println(runtime.NumCPU())

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

	start := time.Now()

	for _, link := range links {
		checkLink(link)
	}

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
}
