package main

import (
	"fmt"
	"time"
)

func main(){
	channel1 := make(chan string)
	channel2 := make(chan string)

	go print_ntimes("channel 1", 5, channel1)
	go print_ntimes("channel 2", 5, channel2)

	// we can close connection to any channel by using line below
	// close(channel2)

	// select is switch type statment in go. select is only applicable for go routines
	for i:= 1 ; i <= 2; i++{
		select {
			case message1 := <- channel1 : fmt.Println(message1)
			case message2 := <- channel2 : fmt.Println(message2)
			// default : fmt.Println("channel is busy")
		}
	}
}

func print_ntimes(name string, n int, c chan <-  string){
	for i := 1; i <= 5; i++ {
		fmt.Println(name," i : ",i)
		time.Sleep(time.Millisecond * 1000)
	}
	c <- name + " finished!"
}