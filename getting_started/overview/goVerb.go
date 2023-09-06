package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func goVerb() {
	fmt.Println("Start main")
	go side() // this is not blocking, main will continue running
	fmt.Println("Return to main")
	time.Sleep(5 * time.Second)
	fmt.Println("End main")
}

func side() {
	fmt.Println("Start side process")
	time.Sleep(1 * time.Second)
	fmt.Println("End side process")
}

func waitForAll() {
	fmt.Println("Start main")
	/* tells the main function how many times wg.Done() function needs to be
	 * called from the side() before it will consider moving past wg.Wait. */
	wg.Add(2)
	go subProcess()
	go subProcess()
	fmt.Println("Return to main")
	wg.Wait()
	fmt.Println("End main")
}

func subProcess() {
	fmt.Println("Start side process")
	time.Sleep(1)
	fmt.Println("End side process")
	wg.Done()
}
