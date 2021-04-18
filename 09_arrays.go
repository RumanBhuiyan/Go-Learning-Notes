package main

import "fmt"

func main(){
	// declaring arrays
   var numbers [5]int 
   var another_set =[5] int {1,2,3,4,5}

   printElements(another_set,5)

   // changing items of numbers array
   var i int 
   for i = 0 ; i < len(numbers) ; i++ {
	   numbers[i] = (i+1)*5
   }  
   
   printElements(numbers,len(numbers))
}

// decalaring a void function
func printElements(numbers [5]int,size int){
	var i int 
	for i = 0 ; i < size ; i++ {
      fmt.Println(i+1,"element: ",numbers[i])
	} 
}
