package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	wg              sync.WaitGroup
	widgetInventory int32 = 1000
)

func main() {
	fmt.Println("Starting inventory count =", widgetInventory)
	wg.Add(2)
	go makeSales()
	go newPurchases()
	wg.Wait()
	fmt.Println("Ending inventory count =", widgetInventory)
}

func makeSales() {
	for i := 0; i < 300000; i++ {
		atomic.AddInt32(&widgetInventory, -100)
	}
	wg.Done()
}

func newPurchases() {
	for i := 0; i < 300000; i++ {
		atomic.AddInt32(&widgetInventory, 100)
	}
	wg.Done()
}
