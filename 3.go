package main

import (
	"fmt"
	"sync"
)

var b = make(chan int, 2)
var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go Work("goroutine1")
	go Work("goroutine2")
	go Work("goroutine3")
	wg.Wait()
	fmt.Println("successful")
}
func Work(workName string) {
	b <- 5
	fmt.Println(workName)
	<-b
	wg.Done()
}
