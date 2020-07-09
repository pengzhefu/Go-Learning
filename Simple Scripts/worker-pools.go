// Our running program shows the 5 jobs being executed by various workers. The program only takes about 2 seconds despite doing about 5 seconds of total work because there are 3 workers operating concurrently.

package main

import (
    "fmt"
	"time"
	"sync"
)
// This is the function we’ll run in every goroutine. Note that a WaitGroup must be passed to functions by pointer.
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()  // here will execute when 
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(5 * time.Second)
        fmt.Println("worker", id, "finished job", j)
        // results <- j * 2  // results represents that 
    }
}

func main() {

	fmt.Println("start the job")

	var wg sync.WaitGroup
    const numJobs = 5
    jobs := make(chan int, numJobs)
    // results := make(chan int, numJobs)

	// start three workers
    for w := 1; w <= 3; w++ {
		wg.Add(1)  // Launch several goroutines and increment the WaitGroup counter for each.
        go worker(w, jobs, &wg)
	}
	
	// inputting data into jobs
	fmt.Println("inputting data")

    for j := 1; j <= numJobs; j++ {
		jobs <- j
		fmt.Println("here to demonstarte that when a channel has content the worker will start")
	}
    close(jobs)  // if there is goroutine to use channel, not have to close

    // for a := 1; a <= numJobs; a++ {
    //     <-results  // blocking channel, wait till there is content in channel
	// }
	wg.Wait()  // to wait for the finish of each goroutine
	
}