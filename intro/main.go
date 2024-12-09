package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	s := []string{"queen", "ochieng", "mama", "aiyana"}
	for _, str := range s {
		go Printstr(str)
	}

}
func Printstr(str string) {
	fmt.Println("this is goroutine", str)
	time.Sleep(time.Second)
}
