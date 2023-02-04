package main

import (
	"fmt"
	"sync"
)

var (
	wg              sync.WaitGroup
	mutex                 = sync.Mutex{}
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
		mutex.Lock()
		widgetInventory -= 100
		mutex.Unlock()
	}
	wg.Done()
}

func newPurchases() {
	for i := 0; i < 300000; i++ {
		mutex.Lock()
		widgetInventory += 100
		mutex.Unlock()
	}
	wg.Done()
}
