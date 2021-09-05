package main

import (
	"fmt"
	"net/http"
)

func main(){

	response , error := http.Get("https://fakestoreapi.com/products")

	if (error != nil){
		fmt.Println("Error : ",error)
	}else{
		fmt.Println(response)
	}
}