package main

import (
	"fmt"
	"time"
)

func main(){
	// go function_name will run that function in a different channel parallely that is called
	// go routine
	go print_ntimes("channel 1", 5)
	fmt.Println("main channel is running...")

	// taking input we are forcing main channel to wait until other channels finishes execution
	var number int
	fmt.Scan(&number)
}

func print_ntimes(name string ,n int){
	for i := 1; i <= n ; i++ {
		fmt.Println(name," i : ",i)
		time.Sleep(time.Millisecond * 1000)
	}
}