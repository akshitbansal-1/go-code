package main

import "fmt"
func main() {
	c := make(chan int)

	go func() {
		x := <- c
		fmt.Println(x)
	}()

	c <- 10
	// c <- 20
	// x := <-c
	// fmt.Println(x)
	fmt.Scanln()
}