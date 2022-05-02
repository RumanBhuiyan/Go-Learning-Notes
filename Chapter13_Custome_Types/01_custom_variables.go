package main

import (
	"fmt"
)

type kelvin float64
type celcius float64

func main(){
	var total_celcius celcius
	var total_kelvin kelvin

	c := 27.04
	k := 312.24

	//total_celcius += c; // type mismatch celcius and float64
	total_celcius += celcius(c)
	total_kelvin += kelvin(k)

	fmt.Printf("Type of c : %T value : %[1]v \n",total_celcius)
	fmt.Printf("Type of k : %T value : %[1]v \n",total_kelvin)
}