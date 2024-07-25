package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf(" %d -> Task [ %s ] is running \n", i, name)
		time.Sleep(time.Second)
	}
}
func taskComWaitGroup(name string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Printf(" %d -> Task [ %s ] is running \n", i, name)
		time.Sleep(time.Second)
		wg.Done()
	}
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(15)
	go taskComWaitGroup("routine 1", &waitGroup)
	go taskComWaitGroup("routine 2", &waitGroup)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf(" %d -> Task [ %s ] is running \n", i, "routine 3")
			time.Sleep(time.Second)
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()

}

// thread
func goRoutines() {
	go task("routine 1")
	go task("routine 2")

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf(" %d -> Task [ %s ] is running \n", i, "routine 3")
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("Main routine is done")

}

func waitGroup() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(15)
	go taskComWaitGroup("routine 1", &waitGroup)
	go taskComWaitGroup("routine 2", &waitGroup)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf(" %d -> Task [ %s ] is running \n", i, "routine 3")
			time.Sleep(time.Second)
		}
	}()

	waitGroup.Wait()
	//fmt.Println("Main routine is done")

}
