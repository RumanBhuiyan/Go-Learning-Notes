package main

import (
	"fmt"
	"time"
)

func main(){
	// creating a channel which will send string value
	channel1 := make(chan string)

	fmt.Println("main channel started...")

	// running subroutines concurrently
	go sender_channel("channel 1", 5, channel1)
	go receiver_channel("channel 2", 5, channel1)

	// forcing main channel to wait until two channel finishes execution
	var number int
	fmt.Scan(&number)
}

// c chan <- string , this type of channel can only send values to channel 
func sender_channel(name string, n int , c chan <- string){
	for i:= 1; i <= 5; i++ {
		fmt.Println(name," i : ",i)
		time.Sleep(time.Millisecond * 1000)
	}
	// sending value via the channel
	c <- name + " finished"
}

// c <- chan string , this type of channel can only receives values from channel 
func receiver_channel(name string, n int , c <- chan string){
	for i:= 1; i <= 5; i++ {
		fmt.Println(name," i : ",i)
		time.Sleep(time.Millisecond * 1000)
	}
	// receiving value via the channel
	message := <- c
	fmt.Println(message)
	fmt.Println(name," finished!")
}