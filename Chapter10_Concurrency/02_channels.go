package main

import (
	"fmt"
	"time"
)

func main(){
	// creating a channel which will send string value
	channel1 := make(chan string)

	// running subroutines concurrently
	go print_ntimes("channel 1", 5, channel1)
	go print_ntimes("channel 2", 5, channel1)

	fmt.Println("main channel is running...")

	// wait for the messages from different channels
	for i := 1 ; i <= 2; i++ {
		message := <- channel1
		fmt.Println("message : ",message)
	}
	fmt.Println("main channel finished!")
}
// c chan string are called bidirectional channel because they can send value to channel
// and receive value from channel
func print_ntimes(name string ,n int, c chan string){
	for i := 1; i <= n ; i++ {
		fmt.Println(name," i : ",i)
		time.Sleep(time.Millisecond * 1000)
	}
	c <- name+" finished"
}