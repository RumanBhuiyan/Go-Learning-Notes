package main

import "fmt"
import "errors"

func main(){
  var result , error = getResult(5,0)

  if (error == nil){
	  fmt.Println(result)
  }else {
	  fmt.Println(error)
  }
}

func getResult(a,b int) (float32,error) {
    if ( b == 0){
		return 0,errors.New("can't divide by zero")
	}
	// type casting first 
	return float32(a)/float32(b) ,nil;
}