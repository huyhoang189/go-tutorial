package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for n := range jobs {
		fmt.Println("worker : ", id, "processed task ", n)
		results <- fibo(n)
	}
}

func fibo(n int) int {
	//fmt.Println(n)
	if n <= 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}

func main() {
	fmt.Println("Pattern Work Pool")
	dt_1 := time.Now()

	number := 30
	numberWorkers := 4
	jobs := make(chan int, number)
	results := make(chan int, number)

	for i := 1; i <= numberWorkers; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= number; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 1; i <= number; i++ {
		//fmt.Println(<-results)
		<-results
	}

	//for i := 1; i <= number; i++ {
	//	fmt.Println(fibo(i))
	//}

	dt_2 := time.Now()
	fmt.Println((dt_2.Sub(dt_1)).String())
}
