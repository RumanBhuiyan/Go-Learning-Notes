package main

import (
	"fmt"
	"time"
)

func main(){
	channel1 := make(chan bool)
	channel2 := make(chan int64)

	go sendFibo(channel1, channel2)

	var fibo int64
	for i:= 1 ; i <= 10 ; i++{
		channel1 <- true
		fibo = <- channel2
		fmt.Println(fibo)
	}
}

func sendFibo(channel1 chan bool, channel2 chan int64) {
	var m bool
	var fibo int64

	n1 := 1
	n2 := 1

	m = <- channel1
	if m {
		channel2 <- int64(n1)
		time.Sleep(time.Second * 1)
	}

	m = <- channel1
	if m {
		channel2 <- int64(n2)
		time.Sleep(time.Second * 1)
	}


	for i := 3; i <= 10 ; i++ {
		m = <- channel1
		if m{
			fibo = int64(n1) + int64(n2)
			n1 = n2
			n2 = int(fibo)
			channel2 <- fibo
			time.Sleep(time.Second * 1)
		}
	} 
}