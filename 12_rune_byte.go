package main

import "fmt"

func main(){

	var character rune = '$'
	fmt.Println("ASCII value of $ : ",character)

	// we should store those characters which ASCII value ranges from
	// 0-255 . If any character ASCII value becomes higher than use rune type
	var singleChar  byte= 'a'
	fmt.Println("ASCII Value of a : ",singleChar)
}