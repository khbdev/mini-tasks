package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
	mu sync.Mutex
)


func worker(id int, jobs <- chan int, results chan <- int, wg *sync.WaitGroup){
	defer wg.Done()
	mu.Lock()
	for job := range jobs {
		fmt.Printf("Worker id %d\n", id)
		results <- 2 * job
	}
	mu.Unlock()
}



func main(){
jobs := make(chan int, 10)
result := make(chan int, 10)



for i := 1; i <= 3; i++ {
	wg.Add(1)
go func(id int) {
	worker(id, jobs, result, &wg)
}(i)
}


for i := 1; i < 10; i++ {
	jobs <- i
}
close(jobs)


for i := 1; i < 10; i++ {
	val := <- result
	fmt.Println("result", val)
}
}