/*
	arguments vs parameters : function receives parameters and function is invoked with arguments
	functions vs methods : methods become bind with class but functions dont.
						 methods are accessed using . but functions can be called by their name
	
	functions vs closures : functions can't keep track of variables declared inside it after execution
							but closures can hold reference of variables surrounding it.
*/
package main

import (
	"fmt"
)

type celcius float64
type kelvin float64

// here (k kelvin) is called receiver, a method will have just one receiver but may have multiple parameters
// when we call kel.to_celcius() then value of kel is passed by value to k. Thats why k gets a copy of kel value

func (k kelvin) to_celcius() celcius{
	c := celcius(k - 273)
	return c
}
// here c receives referece of cel as a result it can change actual value of cel
func (c *celcius) increment(){
	*c += celcius(1)
}

func (c celcius) to_kelvin() kelvin {
	k := kelvin(c + 273)
	return k
}

func main(){
	cel := celcius(25)
	kel := kelvin(300) 

	fmt.Printf("c(celcius) : %v c(kelvin) : %v \n",cel, cel.to_kelvin())
	fmt.Printf("kelvin(kelvin) : %v c(celcius) : %v \n",kel, kel.to_celcius())

	cel.increment()
	fmt.Println("value of cel after incrementing : ",cel)
}
