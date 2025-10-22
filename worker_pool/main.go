package main

import (
	"fmt"
	"sync"
	"time"
)

/*
How to handle 1000 concurrent API requests efficiently?

I use goroutines for concurrency,
a worker pool to limit heavy processing,
DB connection pooling,
caching for frequent data,
rate limiting for protection,
horizontal scaling with load balancers.
*/
// generateJobs returns a slice of jobs
func generateJobs(n int) []int {
	jobs := make([]int, n)
	for i := 0; i < n; i++ {
		jobs[i] = i + 1
	}
	return jobs
}

// processJob simulates doing work for a job
func processJob(job int) {
	fmt.Printf("Worker processing job %d\n", job)
	time.Sleep(2 * time.Second) // simulate work
	fmt.Printf("Job %d done\n", job)
}

// worker is a goroutine that processes a job using a semaphore
func worker(job int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()          // notify WaitGroup when done
	defer func() { <-sem }() // release semaphore slot

	processJob(job)
}

func main() {
	jobs := generateJobs(10)      // create 10 jobs
	sem := make(chan struct{}, 6) // limit to 6 concurrent jobs
	var wg sync.WaitGroup

	for _, job := range jobs {
		wg.Add(1)         // increment WaitGroup counter
		sem <- struct{}{} // acquire semaphore slot

		go worker(job, sem, &wg) // start job in a goroutine
	}

	wg.Wait() // wait for all jobs to finish
	fmt.Println("All jobs completed!")
}
