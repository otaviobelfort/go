package main

import "time"

func worker(workerID int, data chan int) {
	for {
		select {
		case value := <-data:
			println("Worker ID: ", workerID, " Value: ", value)
			time.Sleep(time.Second)
		}
	}
	//for value := range data {
	//	println("Worker ID: ", workerID, " Value: ", value)
	//	//time.Sleep(time.Second)
	//}
}
func main() {
	processWorker()
}

func processWorker() {
	data := make(chan int)

	go worker(1, data)
	go worker(2, data)

	for i := 0; i < 10; i++ {
		data <- i
	}
	close(data)

}
