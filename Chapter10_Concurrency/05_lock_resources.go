package main

import (
	"fmt"
	"sync"
	"time"
)

var myLock sync.Mutex
var counter int= 0

func main(){
	channel := make(chan string)

	go incrementByOne(channel)
	go incrementByTwo(channel)

	message1 := <- channel
	message2 := <- channel

	fmt.Println(message1)
	fmt.Println(message2)
}

func incrementByOne(c chan <- string){
	for i := 1; i <= 5; i++ {
		myLock.Lock()
			fmt.Println("IncrementByOne locked resources...")
			counter += 1
			fmt.Printf("incrementByOne counter : %v \n",counter)
		myLock.Unlock()
		fmt.Println("IncrementByOne released lock")
		time.Sleep(time.Second * 1)
	}
	c <- "incrementByOne finished"
}

func incrementByTwo(c chan <- string){
	for i := 1; i <= 5; i++ {
		myLock.Lock()
			fmt.Println("IncrementByTwo locked resources...")
			counter += 2
			fmt.Printf("incrementByTwo counter : %v \n",counter)
		 myLock.Unlock()
		fmt.Println("IncrementByTwo released lock")
		time.Sleep(time.Second * 1)
	}
	c <- "incrementByTwo finished"
}