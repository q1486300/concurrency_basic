package main

import (
	"fmt"
	"sync"
)

var (
	wg              sync.WaitGroup
	mutex                 = sync.Mutex{}
	widgetInventory int32 = 1000
	newPurchase           = sync.NewCond(&mutex)
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
	for i := 0; i < 3000; i++ {
		mutex.Lock()
		for widgetInventory-100 < 0 {
			// condition variable: 避免庫存變成負數
			newPurchase.Wait()
		}
		widgetInventory -= 100
		fmt.Println(widgetInventory)
		mutex.Unlock()
	}
	wg.Done()
}

func newPurchases() {
	for i := 0; i < 3000; i++ {
		mutex.Lock()
		widgetInventory += 100
		fmt.Println(widgetInventory)
		newPurchase.Signal() // 發出信號讓等待的 condition variable 停止阻塞
		mutex.Unlock()
	}
	wg.Done()
}
