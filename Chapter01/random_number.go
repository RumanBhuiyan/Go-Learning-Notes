package main

import (
	"fmt"
	"math/rand"
)

func main(){
	//random number is generated and cached that's why it shows same result again & agian
	// pick a number within range 0-10 in exclusive manner where 10 isn't included
	number := rand.Intn(10)

	fmt.Println("number : ",number)

	number = rand.Intn(10)
	fmt.Println("number : ",number)

	// taking a random number within a range like 10-20
	a := rand.Intn(10) + 10
	fmt.Println("range 10-20 a : ",a)
}