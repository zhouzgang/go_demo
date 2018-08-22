package main

import (
	"time"
	"fmt"
)

func main() {
	go spinner(1000 * time.Microsecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rF(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `_\|/`{
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	//println(x)
	return fib(x-1) + fib(x-2)
}