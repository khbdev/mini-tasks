package main

import (
	"fmt"
	"sync"
)




var (
	mu sync.Mutex
)


func worker(id int, jobs <- chan int, wg *sync.WaitGroup){
	defer wg.Done()

mu.Lock()
	for  job := range jobs {
		fmt.Printf("gorutina id %d, worker id %d\n", id, job)
	}
	mu.Unlock()
}


func main(){
	jobs :=  make(chan int, 10)


   var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
				
			worker(id, jobs, &wg)
		}(i)
	}

	for i := 1; i <= 10; i++ {
	
jobs <- i
	}
	close(jobs)

	wg.Wait()

	

}