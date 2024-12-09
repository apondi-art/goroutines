package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	Prints("Queen")
	fmt.Println(time.Since(start))

}

func Prints(s string) {

	fmt.Printf("Hello ,how are you %s\n", s)

}
