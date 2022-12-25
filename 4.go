package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var aa = make(chan int, 2)

func main() {
	wg.Add(5)
	go func() {
		aa <- 5
		var a = 1
		for ; a < 10; a++ {
			if a == a {
				fmt.Println(a)
			}
		}
		wg.Done()
		<-aa
	}()
	go func() {
		aa <- 5
		var a = 10
		for ; a < 100; a++ {
			b := a % 10
			c := a / 10
			if a == b*b+c*c {
				fmt.Println(a)
			}
		}
		wg.Done()
		<-aa
	}()
	go func() {
		aa <- 5
		var a = 100
		for ; a < 1000; a++ {
			b := a % 10
			c := (a % 100) / 10
			d := a / 100
			if a == b*b*b+d*d*d+c*c*c {
				fmt.Println(a)
			}
		}
		wg.Done()
		<-aa
	}()
	go func() {
		aa <- 5
		var a = 1000
		for ; a < 10000; a++ {
			b := a % 10
			c := (a % 100) / 10
			d := (a % 1000) / 100
			e := a / 1000
			if a == b*b*b*b+d*d*d*d+c*c*c*c+e*e*e*e {
				fmt.Println(a)
			}
		}
		wg.Done()
		<-aa
	}()
	go func() {
		aa <- 5
		var a = 10000
		for ; a < 100000; a++ {
			b := a % 10
			c := (a % 100) / 10
			d := (a % 1000) / 100
			e := (a % 10000) / 1000
			f := a / 10000
			if a == b*b*b*b*b+d*d*d*d*d+c*c*c*c*c+e*e*e*e*e+f*f*f*f*f {
				fmt.Println(a)
			}
		}
		wg.Done()
		<-aa
	}()

	wg.Wait()
}
