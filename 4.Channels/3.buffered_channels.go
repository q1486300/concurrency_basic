package main

import "fmt"

func main() {
	c := make(chan string, 3)
	c <- "Hello "
	c <- "Earth "
	c <- "from Mars "
	//c <- "from Venus" // Deadlock!

	msg := <-c
	fmt.Print(msg)

	msg = <-c
	fmt.Print(msg)

	msg = <-c
	fmt.Println(msg)
}
