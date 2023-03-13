package main

import (
	"fmt"
	"time"
)

// crawl mock 10K URLs với buffered channel và goroutines

func worker(queue <-chan int, workerNum int) {
	for v := range queue {
		fmt.Printf("Worker %d process crawl url %d \n", workerNum, v)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d done crawl url %d \n", workerNum, v)
	}

}

func startQ(maxWorker int, urlsNum int) <-chan int {
	queue := make(chan int, maxWorker)

	go func() {
		for i := 0; i < urlsNum; i++ {
			queue <- i
			fmt.Printf("URL %d has been enqueued\n", i)
		}
		close(queue)
	}()
	return queue
}

func main() {
	var (
		urlsNum     = 10
		maxRoutines = 5
	)
	fmt.Printf("Starting crawl %d Urls with %d goroutines \n", urlsNum, maxRoutines)

	q := startQ(maxRoutines, urlsNum)
	for i := 0; i < maxRoutines; i++ {
		go worker(q, i)
	}

	time.Sleep(time.Minute * 2)
}
