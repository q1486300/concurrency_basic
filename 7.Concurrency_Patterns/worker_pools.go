package main

import (
	"fmt"
	"time"
)

func main() {
	const numJobs = 10
	jobsChan := make(chan int, numJobs)
	completedJobsChan := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobsChan, completedJobsChan)
	}

	for j := 1; j <= numJobs; j++ {
		jobsChan <- j
	}
	close(jobsChan)

	for i := 1; i <= numJobs; i++ {
		<-completedJobsChan
	}
}

func worker(id int, jobsChan <-chan int, completedJobsChan chan<- int) {
	for j := range jobsChan {
		fmt.Printf("worker %d started job %d with %d jobs left to process\n", id, j, len(jobsChan))
		time.Sleep(time.Second * 2)
		fmt.Printf("worker %d    \t\t\t    finished job %d\n", id, j)
		completedJobsChan <- j
	}
}
