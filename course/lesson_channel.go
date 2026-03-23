//go run lesson_channel.go

package main

import "fmt"



func main() {
{
	//42
	buffered_channel := make( chan int, 1)
	buffered_channel <- 42
	fmt.Print(<-buffered_channel)
	//fatal error: all goroutines are asleep - deadlock!
	unbuffered_channel := make( chan int)
	unbuffered_channel <- 42
	fmt.Print(<-unbuffered_channel)
}}