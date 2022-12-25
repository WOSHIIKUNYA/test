package main

import "fmt"

// defer是让程序最后执行该语句，但是当读到17行时程序就停止了就还来不急进行输出，如果想要输出就要在17行后面加语句，总之defer不能放在程序执行最后。
func main() {
	var a = true
	defer func() {
		fmt.Println("1")
	}()
	if a {
		fmt.Println("2")
		return
	}
	defer func() {
		fmt.Println("3")
	}()
}
