package main

import "fmt"

func main(){
	// declaring maps 
	var keys map[string]int
	// creating map
    keys  = make(map[string]int)

	keys["apple"]= 1
	keys["mango"]= 2
	keys["orange"]= 3
	keys["jam"]= 4
	keys["grape"]= 5
    
	// before deleting 
	showMapItems(keys)
     fmt.Println("------------")
	// after deleting item 
    delete(keys,"grape")
    showMapItems(keys)
}
func showMapItems(myMap map[string]int){
    var i string ;
	for i = range myMap {
		fmt.Println(i," : ",myMap[i])
	}
}