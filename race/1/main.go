package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup
var result int

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	waitGroup.Add(2)
	go runProcess("P1", 20)
	go runProcess("P2", 20)
	waitGroup.Wait()
	fmt.Println("Final result:", result)
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		z := result
		z++

		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		result = z
		fmt.Println(name, "->", i, "Partial result: ", result)
	}
	waitGroup.Done()
}