package main

import (
	"fmt"
	"errors"
)

func main(){
	err := errors.New("error : testing error")
	fmt.Println(err)
}