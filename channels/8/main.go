package main

import "fmt"

func main() {
	// fan out
	c := generate(5, 10, 15, 25)
	d1 := divide(c)
	d2 := divide(c)

	fmt.Println("D1", <-d1)
	fmt.Println("D2", <-d2)
	fmt.Println("D1", <-d1)
	fmt.Println("D2", <-d2)

}

func generate(numbers ...int) chan int {
	channel := make(chan int)
	go func() {
		for _, number := range numbers {
			channel <- number
		}
		close(channel)
	}()

	return channel
}

func divide(input chan int) chan int  {
	channel := make(chan int)
	go func() {
		for number := range input {
			channel <- number / 5
		}
		close(channel)
	}()

	return channel
}