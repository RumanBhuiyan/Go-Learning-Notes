package main

import (
	"fmt"
)

type celcius float64
type kelvin float64
type kelvinTemp  func() kelvin

func kelvin_room_temparature() kelvin {
	return kelvin(0 + 273)
}

func celcius_to_kelvin(c celcius, fn kelvinTemp) kelvin{
	if c == 0{
		return fn()
	}else {
		return kelvin(273 + float64(c))
	}
}

func main(){
	var c1 celcius = 0
	var c2 celcius = 20

	fmt.Printf("c1 (celcius) : %v c1(kelvin) : %v \n",c1,celcius_to_kelvin(c1, kelvin_room_temparature))
	fmt.Printf("c2 (celcius) : %v c2(kelvin) : %v \n",c2,celcius_to_kelvin(c2, kelvin_room_temparature))

}

