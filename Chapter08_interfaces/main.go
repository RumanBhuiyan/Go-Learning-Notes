package main

import (
	"fmt"
)

// defining interface : type <interface_name> interface
type Natures interface{
	eat()
	speak()
}

type Animal struct{
	name string
	nature Natures // Animal type struct instances has eat(), speak()
}

/* There is no keyword like implements, extends in Go to implement methods of interfaces
 following way is the only way to get the job done. N.B => We wont get any compile error 
 or run time error if we dont implement interface methods */
func (animal *Animal) eat(){
	fmt.Println(animal.name," eats insects")
}

func (animal *Animal) speak(){
	fmt.Println(animal.name," speaks kooh! kooh!")
}

func main(){
	bird := Animal{
		name : "Cuckoo",
	}

	fmt.Println("bird.name : ", bird.name)
	bird.eat()
	bird.speak()
}