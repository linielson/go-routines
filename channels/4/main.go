package main

import "fmt"

func main() {
	channel := make(chan int)
	//channel <- 10 // fatal error: all goroutines are asleep - deadlock!
	go func() {
		channel <- 10
	}()
	fmt.Println(<-channel)
}
