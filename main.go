package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var result int
	var count int
	processors := runtime.GOMAXPROCS(0)
	fmt.Println(processors)
	for i := 0; i < 8000; i++ {
		count++
		fmt.Println(count)
		go func() {
			for {
				result++
				//time.Sleep(0)
			}
		}()
	}
	time.Sleep(time.Second) //wait for go function to increment the value.
	fmt.Println("result =", result)
}
