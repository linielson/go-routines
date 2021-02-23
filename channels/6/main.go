package main

import "fmt"

//Pipeline
func main() {
	numbers := generate(2, 4, 6) // generate and populate a channel ONE
	result := divide(numbers)

	fmt.Println(<-result) // remove item from channel TWO
	fmt.Println(<-result) // remove item from channel TWO
	fmt.Println(<-result) // remove item from channel TWO
}

func generate(numbers ...int) chan int {
	channel := make(chan int) // create channel ONE
	go func() {
		for _, number := range numbers {
			channel <- number
		}
	}()
	return channel
}

func divide(input chan int) chan int  {
	channel := make(chan int) // create channel TWO
	go func() {
		for number := range input { // remove item from channel ONE
			channel <- number / 2 // add item at channel TWO
		}
		close(channel)
	}()

	return channel
}