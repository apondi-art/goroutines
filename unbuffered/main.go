package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	workingNum := 1
	wg.Add(workingNum)
	start := time.Now()

	go func() {

		
		PrintS("Hello", &wg)

	}()
	

	// time.Sleep(2 * time.Millisecond)

	go func() {
		
		PrintS("how", &wg)
	}()
	wg.Wait()
	s := time.Since(start)
	fmt.Println(s)
}

func PrintS(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Mr Obima %s\n", s)
}
